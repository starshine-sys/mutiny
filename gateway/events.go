package gateway

import (
	"github.com/starshine-sys/mutiny/revolt"
)

// DisconnectEvent is emitted when the gateway connection is closed.
type DisconnectEvent struct{}

// UnknownEvent is emitted when an unknown event is received.
type UnknownEvent struct {
	Type    string
	RawData []byte
}

// ReadyEvent ...
type ReadyEvent struct {
	Users    []revolt.User    `json:"users"`
	Servers  []revolt.Server  `json:"servers"`
	Channels []revolt.Channel `json:"channels"`
}

// MessageEvent ...
type MessageEvent struct {
	revolt.Message
}

// MessageUpdateEvent ...
type MessageUpdateEvent struct {
	ID   string         `json:"id"`
	Data revolt.Message `json:"data"`
}

// MessageDeleteEvent ...
type MessageDeleteEvent struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
}

// ChannelCreateEvent ...
type ChannelCreateEvent struct {
	revolt.Channel
}

// ChannelUpdateEvent ...
type ChannelUpdateEvent struct {
	ID    string                 `json:"id"`
	Data  ChannelUpdateEventData `json:"data"`
	Clear string                 `json:"clear,omitempty"`
}

// Update updates the given channel with data from the event.
func (ev ChannelUpdateEvent) Update(ch *revolt.Channel) {
	if ev.Data.Name != nil {
		ch.Name = *ev.Data.Name
	}
	if ev.Data.Type != nil {
		ch.Type = *ev.Data.Type
	}
	if ev.Data.NSFW != nil {
		ch.NSFW = *ev.Data.NSFW
	}
	if ev.Data.LastMessageID != nil {
		ch.LastMessageID = *ev.Data.LastMessageID
	}
	if ev.Data.Permissions != nil {
		ch.Permissions = *ev.Data.Permissions
	}
	if ev.Data.DefaultPermissions != nil {
		ch.DefaultPermissions = *ev.Data.DefaultPermissions
	}
	if ev.Data.RolePermissions != nil {
		ch.RolePermissions = ev.Data.RolePermissions
	}
	if ev.Data.DMActive != nil {
		ch.DMActive = *ev.Data.DMActive
	}
	if ev.Data.GroupOwnerID != nil {
		ch.GroupOwnerID = *ev.Data.GroupOwnerID
	}
	if ev.Data.Recipients != nil {
		ch.Recipients = *ev.Data.Recipients
	}

	switch ev.Clear {
	case "Icon":
		ch.Icon = nil
	case "Description":
		ch.Description = ""
	}
}

// ChannelUpdateEventData ...
type ChannelUpdateEventData struct {
	Type *string            `json:"channel_type,omitempty"`
	Name *string            `json:"name,omitempty"`
	Icon *revolt.Attachment `json:"icon,omitempty"`
	NSFW *bool              `json:"nsfw,omitempty"`

	Description *string `json:"description,omitempty,omitempty"`

	LastMessageID *string `json:"last_message_id,omitempty"`

	// Only in groups
	Permissions *uint64 `json:"permissions,omitempty"`
	// Only in servers
	DefaultPermissions *uint64           `json:"default_permissions,omitempty"`
	RolePermissions    map[string]uint64 `json:"role_permissions,omitempty"`

	// If a DM, whether it's active
	DMActive *bool `json:"active,omitempty"`
	// If a group, the owner's ID
	GroupOwnerID *string   `json:"owner,omitempty"`
	Recipients   *[]string `json:"recipients,omitempty"`
}

// ChannelStartTypingEvent ...
type ChannelStartTypingEvent struct {
	ChannelID string `json:"channel_id"`
	UserID    string `json:"user"`
}

// ChannelStopTypingEvent ...
type ChannelStopTypingEvent struct {
	ChannelID string `json:"channel_id"`
	UserID    string `json:"user"`
}
