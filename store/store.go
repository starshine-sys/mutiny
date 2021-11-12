package store

import (
	"emperror.dev/errors"
	"github.com/starshine-sys/mutiny/revolt"
)

// Store is a cache interface.
type Store interface {
	ChannelStore
	UserStore
	ServerStore
}

// ChannelStore is a channel store interface.
type ChannelStore interface {
	Channel(id string) (*revolt.Channel, error)
	PutChannel(*revolt.Channel) error
	RemoveChannel(id string) error
	ResetChannels() error
}

// UserStore is a user store interface.
type UserStore interface {
	User(id string) (*revolt.User, error)
	PutUser(*revolt.User) error
	RemoveUser(id string) error
	ResetUsers() error
}

// ServerStore is a server store interface.
type ServerStore interface {
	Server(id string) (*revolt.Server, error)
	PutServer(*revolt.Server) error
	RemoveServer(id string) error
	ResetServers() error
}

// ErrNotFound ...
const ErrNotFound = errors.Sentinel("not found in store")
