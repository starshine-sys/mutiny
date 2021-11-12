package api

import "github.com/starshine-sys/mutiny/revolt"

// User returns the given user.
func (c *Client) User(id string) (*revolt.User, error) {
	var u revolt.User
	err := c.RequestJSON(&u, "GET", EndpointUsers+id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
