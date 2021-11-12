package defaultstore

import (
	"sync"

	"github.com/starshine-sys/mutiny/revolt"
	"github.com/starshine-sys/mutiny/store"
)

// DefaultChannelStore ...
type DefaultChannelStore struct {
	sync.RWMutex

	chans map[string]*revolt.Channel
}

var _ store.ChannelStore = (*DefaultChannelStore)(nil)

// Channel returns a channel from the store.
func (s *DefaultChannelStore) Channel(id string) (*revolt.Channel, error) {
	s.RLock()
	defer s.RUnlock()

	ch, ok := s.chans[id]
	if ok {
		return ch, nil
	}

	return nil, store.ErrNotFound
}

// PutChannel puts a channel into the store.
func (s *DefaultChannelStore) PutChannel(ch *revolt.Channel) error {
	s.Lock()
	s.chans[ch.ID] = ch
	s.Unlock()
	return nil
}

// RemoveChannel removes a channel from the store.
func (s *DefaultChannelStore) RemoveChannel(id string) error {
	s.Lock()
	delete(s.chans, id)
	s.Unlock()
	return nil
}

// ResetChannels resets the channel store.
func (s *DefaultChannelStore) ResetChannels() error {
	s.Lock()
	s.chans = map[string]*revolt.Channel{}
	s.Unlock()
	return nil
}
