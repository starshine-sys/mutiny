package revolt

// Channel ...
type Channel struct {
	ID       string      `json:"_id"`
	ServerID string      `json:"server,omitempty"`
	Type     string      `json:"channel_type"`
	Name     string      `json:"name"`
	Icon     *Attachment `json:"icon"`
	NSFW     bool        `json:"nsfw"`

	Description string `json:"description,omitempty"`

	LastMessageID string `json:"last_message_id"`

	// Only in groups
	Permissions uint64 `json:"permissions"`
	// Only in servers
	DefaultPermissions uint64            `json:"default_permissions"`
	RolePermissions    map[string]uint64 `json:"role_permissions"`

	// If a DM, whether it's active
	DMActive bool `json:"active"`
	// If a group, the owner's ID
	GroupOwnerID string   `json:"owner"`
	Recipients   []string `json:"recipients"`
}
