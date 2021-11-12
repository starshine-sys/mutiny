package gateway

import "log"

func (g *Gateway) handleReady(ev *ReadyEvent) {
	log.Printf("Received ready event with %v channels, %v servers, %v users", len(ev.Channels), len(ev.Servers), len(ev.Users))

	for _, ch := range ev.Channels {
		g.Store.PutChannel(&ch)
	}

	for _, s := range ev.Servers {
		g.Store.PutServer(&s)
	}

	for _, u := range ev.Users {
		g.Store.PutUser(&u)
	}
}
