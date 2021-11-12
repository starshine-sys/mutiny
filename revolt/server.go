package revolt

// Server ...
type Server struct {
	ID      string `json:"_id"`
	OwnerID string `json:"owner"`

	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	Icon   *Attachment `json:"icon,omitempty"`
	Banner *Attachment `json:"banner,omitempty"`

	Channels   []string         `json:"channels"`
	Categories []ServerCategory `json:"categories"`

	DefaultPermissions [2]uint64 `json:"default_permissions"`
}

// ServerCategory ...
type ServerCategory struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Channels []string `json:"channels"`
}

// ServerSystemMessages ...
type ServerSystemMessages struct {
	UserJoined string `json:"user_joined"`
	UserLeft   string `json:"user_left"`
	UserKicked string `json:"user_kicked"`
	UserBanned string `json:"user_banned"`
}

// Role ...
type Role struct {
	Name        string    `json:"name"`
	Colour      string    `json:"colour"`
	Permissions [2]uint64 `json:"permissions"`
	Hoist       bool      `json:"hoist"`
	Rank        uint      `json:"rank"`
}
