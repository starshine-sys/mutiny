package gateway

import "github.com/starshine-sys/mutiny/revolt"

// DisconnectEvent is emitted when the gateway connection is closed.
type DisconnectEvent struct{}

// UnknownEvent is emitted when an unknown event is received.
type UnknownEvent struct {
	Type    string
	RawData []byte
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
