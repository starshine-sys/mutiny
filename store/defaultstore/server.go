package defaultstore

import (
	"sync"

	"github.com/starshine-sys/mutiny/revolt"
	"github.com/starshine-sys/mutiny/store"
)

// DefaultServerStore ...
type DefaultServerStore struct {
	sync.RWMutex

	servers map[string]*revolt.Server
}

var _ store.ServerStore = (*DefaultServerStore)(nil)

// Server returns a Server from the store.
func (s *DefaultServerStore) Server(id string) (*revolt.Server, error) {
	s.RLock()
	defer s.RUnlock()

	ch, ok := s.servers[id]
	if ok {
		return ch, nil
	}

	return nil, store.ErrNotFound
}

// PutServer puts a Server into the store.
func (s *DefaultServerStore) PutServer(srv *revolt.Server) error {
	s.Lock()
	s.servers[srv.ID] = srv
	s.Unlock()
	return nil
}

// RemoveServer removes a Server from the store.
func (s *DefaultServerStore) RemoveServer(id string) error {
	s.Lock()
	delete(s.servers, id)
	s.Unlock()
	return nil
}

// ResetServers resets the Server store.
func (s *DefaultServerStore) ResetServers() error {
	s.Lock()
	s.servers = map[string]*revolt.Server{}
	s.Unlock()
	return nil
}
