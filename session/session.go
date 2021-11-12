package session

import (
	"github.com/starshine-sys/mutiny/api"
	"github.com/starshine-sys/mutiny/gateway"
)

// Session combines the Gateway and API client into one.
type Session struct {
	*gateway.Gateway
	*api.Client
}

// New creates a new Session.
func New(token string) (*Session, error) {
	s := &Session{}

	s.Client = api.New(token)

	gw, err := s.Client.QueryNode()
	if err != nil {
		return nil, err
	}
	s.Gateway = gateway.New(gw.Websocket, token)

	return s, nil
}
