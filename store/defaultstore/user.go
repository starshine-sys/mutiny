package defaultstore

import (
	"sync"

	"github.com/starshine-sys/mutiny/revolt"
	"github.com/starshine-sys/mutiny/store"
)

// DefaultUserStore ...
type DefaultUserStore struct {
	sync.RWMutex

	users map[string]*revolt.User
}

var _ store.UserStore = (*DefaultUserStore)(nil)

// User returns a User from the store.
func (s *DefaultUserStore) User(id string) (*revolt.User, error) {
	s.RLock()
	defer s.RUnlock()

	ch, ok := s.users[id]
	if ok {
		return ch, nil
	}

	return nil, store.ErrNotFound
}

// PutUser puts a User into the store.
func (s *DefaultUserStore) PutUser(u *revolt.User) error {
	s.Lock()
	s.users[u.ID] = u
	s.Unlock()
	return nil
}

// RemoveUser removes a User from the store.
func (s *DefaultUserStore) RemoveUser(id string) error {
	s.Lock()
	delete(s.users, id)
	s.Unlock()
	return nil
}

// ResetUsers resets the User store.
func (s *DefaultUserStore) ResetUsers() error {
	s.Lock()
	s.users = map[string]*revolt.User{}
	s.Unlock()
	return nil
}
