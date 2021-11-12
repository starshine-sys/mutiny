// Package gateway contains a gateway client and event handlers.
// Sadly, thanks to how Revolt sends events, gateway clients also need a cache,
// so it doesn't work completely standalone.
package gateway

// Much of this code is taken from discordgo: https://github.com/bwmarrin/discordgo/blob/007bf767904825eb3d881daa136698771d3ee0cb/wsapi.go

import (
	"encoding/json"
	"math/rand"
	"sync"
	"time"

	"emperror.dev/errors"
	"github.com/gorilla/websocket"

	"github.com/starshine-sys/mutiny/store"
	"github.com/starshine-sys/mutiny/store/defaultstore"
)

// ErrWSAlreadyOpen is returned when you attempt to open
// a websocket that already is open.
const ErrWSAlreadyOpen = errors.Sentinel("websocket already opened")

// WSDebug is called for websocket debug logging.
var WSDebug = func(tmpl string, args ...interface{}) {
	return
}

// StoreDebug is called for store debug logging.
var StoreDebug = func(tmpl string, args ...interface{}) {
	return
}

// Gateway is a gateway client.
type Gateway struct {
	url, token string
	conn       *websocket.Conn
	listening  chan interface{}

	Handler *Handler
	Store   store.Store

	LastPing, LastPong time.Time

	// Documentation says "every 10 to 30 seconds":
	// https://developers.revolt.chat/websockets/establishing
	// so this is set to a random interval when creating the gateway
	pingRate time.Duration
	// if the time since the last pong is larger than this, trigger a gateway reconnect
	// default: 3 minutes
	MaximumPong time.Duration

	sync.RWMutex
}

// New creates a new Gateway.
// Note: only bot tokens are supported.
func New(url, token string) *Gateway {
	r := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(20)

	g := &Gateway{
		url:         url,
		token:       token,
		Handler:     NewHandler(),
		Store:       defaultstore.New(),
		pingRate:    time.Duration(10+r) * time.Second,
		MaximumPong: 3 * time.Minute,
	}

	// yes i'm using the Handler for internal methods, might as well
	g.Handler.AddHandler(g.handlePong)
	g.Handler.AddHandler(g.handleStoreEvent)

	return g
}

// Open opens the gateway connection.
func (g *Gateway) Open() (err error) {
	// lock the gateway
	g.Lock()
	defer g.Unlock()

	if g.conn != nil {
		return ErrWSAlreadyOpen
	}

	WSDebug("Connecting to gateway at %v", g.url)

	g.conn, _, err = websocket.DefaultDialer.Dial(g.url+"?format=json", nil)
	if err != nil {
		g.conn = nil
		return err
	}

	WSDebug("Connected to gateway")

	g.conn.SetCloseHandler(func(code int, text string) error {
		return nil
	})

	defer func() {
		if err != nil {
			g.conn.Close()
			g.conn = nil
		}
	}()

	auth := struct {
		Type  string `json:"type"`
		Token string `json:"token"`
	}{Type: "Authenticate", Token: g.token}

	WSDebug("Sending authenticate message")

	err = g.conn.WriteJSON(auth)
	if err != nil {
		return err
	}

	_, b, err := g.conn.ReadMessage()
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, &auth); err != nil {
		return err
	}

	if auth.Type != "Authenticated" {
		var e wserror
		if err = json.Unmarshal(b, &e); err != nil {
			return ErrNotAuthenticatedEvent
		}

		if e.Type != "Error" {
			return ErrNotAuthenticatedEvent
		}

		switch e.Error {
		case "LabelMe":
			return ErrUnlabeled
		case "InternalError":
			return ErrInternalError
		case "InvalidSession":
			return ErrInvalidSession
		case "OnboardingNotFinished":
			return ErrOnboardingNotFinished
		case "AlreadyAuthenticated":
			return ErrAlreadyAuthenticated
		default:
			return ErrNotAuthenticatedEvent
		}
	}

	g.LastPong = time.Now().UTC()

	WSDebug("Successfully authenticated")

	g.listening = make(chan interface{})

	go g.heartbeat(g.conn, g.listening)
	go g.listen(g.conn, g.listening)

	return nil
}

// Ping returns the latency between the last ping and pong.
func (g *Gateway) Ping() time.Duration {
	return g.LastPong.Sub(g.LastPing)
}

func (g *Gateway) heartbeat(conn *websocket.Conn, listening <-chan interface{}) {
	ticker := time.NewTicker(g.pingRate)
	defer ticker.Stop()

	WSDebug("Ping interval is %v", g.pingRate)

	for {
		g.RLock()
		last := g.LastPong
		g.RUnlock()

		dat := time.Now().UnixNano()

		WSDebug("Sending ping, data: %v", dat)
		err := g.conn.WriteJSON(ping{
			Type: "Ping",
			Data: dat,
		})
		if err != nil || time.Since(last) > g.MaximumPong {
			if err != nil {
				WSDebug("Error sending ping: %v", err)
			} else {
				WSDebug("Haven't gotten a pong in %v, triggering reconnect", time.Since(last))
			}

			g.Close()
			g.Open()
			return
		}

		select {
		case <-ticker.C:
		case <-listening:
			return
		}
	}
}

func (g *Gateway) listen(conn *websocket.Conn, listening <-chan interface{}) {
	WSDebug("Starting event loop")

	for {
		t, m, err := conn.ReadMessage()
		if err != nil {
			g.RLock()
			same := conn == g.conn
			g.RUnlock()

			if same {
				WSDebug("Error reading from websocket: %v", err)
				err := g.Close()
				if err != nil {
					WSDebug("Error closing gateway: %v", err)
				}

				// reopen gateway
				WSDebug("Reopening gateway")
				g.Open()
			}

			return
		}

		select {
		case <-listening:
			return
		default:
			go g.handleEvent(t, m)
		}
	}
}

func (g *Gateway) handlePong(p *ping) {
	t := time.Unix(0, p.Data)

	WSDebug("Received pong event, data: %v / ping: %v", p.Data, time.Since(t))

	g.Lock()
	g.LastPing = time.Unix(0, p.Data)
	g.LastPong = time.Now().UTC()
	g.Unlock()
}

type event struct {
	Type string `json:"type"`
}

func (g *Gateway) handleEvent(t int, m []byte) {
	// get event type
	var e event
	err := json.Unmarshal(m, &e)
	if err != nil {
		WSDebug("Error unmarshaling event: %v", err)
		return
	}

	fn, ok := eventCreators[e.Type]
	if !ok {
		WSDebug("Received unknown event %s: %s", e.Type, m)
		go g.Handler.Call(&UnknownEvent{
			Type:    e.Type,
			RawData: m,
		})
		return
	}

	ptr := fn()

	if err = json.Unmarshal(m, ptr); err != nil {
		WSDebug("Error unmarshaling event %s: %v", e.Type, err)
	}
	go g.Handler.Call(ptr)
}

// Close closes the gateway connection.
func (g *Gateway) Close() (err error) {
	g.Lock()

	if g.listening != nil {
		WSDebug("Closing listening channel")
		close(g.listening)
		g.listening = nil
	}

	if g.conn != nil {
		err = g.conn.Close()
		if err != nil {
			WSDebug("Error closing websocket: %v", err)
		}
		g.conn = nil
	}

	g.Unlock()

	WSDebug("Emitting disconnect event")
	g.Handler.Call(&DisconnectEvent{})

	return
}
