package session

import "github.com/starshine-sys/mutiny/revolt"

// User returns the given user.
// If the user is not found in the cache, it queries the API.
func (s *Session) User(id string) (*revolt.User, error) {
	u, err := s.Store.User(id)
	if err == nil {
		return u, nil
	}

	u, err = s.Client.User(id)
	if err == nil {
		err = s.Store.PutUser(u)
		return u, err
	}
	return nil, err
}
