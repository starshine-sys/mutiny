package api

import (
	"github.com/google/uuid"
	"github.com/starshine-sys/mutiny/revolt"
	"github.com/starshine-sys/mutiny/util/httputil"
)

// SendMessageData is the data used in c.SendMessage
type SendMessageData struct {
	Content     string         `json:"content"`
	Attachments []string       `json:"attachments,omitempty"`
	Replies     []MessageReply `json:"replies,omitempty"`

	// If not set, will be set to a random UUID.
	Nonce string `json:"nonce"`
}

// MessageReply ...
type MessageReply struct {
	// The message ID to reply to.
	ID string `json:"id"`
	// Whether or not this reply will mention the replied-to user.
	Mention bool `json:"mention"`
}

// SendMessageComplex sends a message to the target channel. Returns a message object on success.
func (c *Client) SendMessageComplex(channelID string, data SendMessageData) (*revolt.Message, error) {
	if data.Nonce == "" {
		data.Nonce = uuid.New().String()
	}

	var m revolt.Message
	err := c.RequestJSON(&m, "POST", EndpointChannels+channelID+EndpointMessages, httputil.WithJSONBody(data))
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// SendMessage sends a text message.
func (c *Client) SendMessage(channelID string, content string) (*revolt.Message, error) {
	return c.SendMessageComplex(channelID, SendMessageData{
		Content: content,
	})
}

// EditMessage edits the given message's content.
func (c *Client) EditMessage(channelID, msgID string, content string) error {
	dat := struct {
		Content string `json:"content"`
	}{Content: content}

	err := c.RequestJSON(nil, "PATCH", EndpointChannels+channelID+EndpointMessages+msgID, httputil.WithJSONBody(dat))
	return err
}
