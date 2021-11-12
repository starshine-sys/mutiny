package defaultstore

import "github.com/starshine-sys/mutiny/revolt"

// DefaultStore is the default, in-memory store.
type DefaultStore struct {
	*DefaultChannelStore
	*DefaultUserStore
	*DefaultServerStore
}

// New returns a new DefaultStore.
func New() *DefaultStore {
	return &DefaultStore{
		DefaultChannelStore: &DefaultChannelStore{
			chans: map[string]*revolt.Channel{},
		},
		DefaultUserStore: &DefaultUserStore{
			users: map[string]*revolt.User{},
		},
		DefaultServerStore: &DefaultServerStore{
			servers: map[string]*revolt.Server{},
		},
	}
}
