package revolt

import "github.com/starshine-sys/mutiny/util/jsonutil"

// Message is a Revolt message.
type Message struct {
	ID        string `json:"_id"`
	Nonce     string `json:"nonce,omitempty"`
	ChannelID string `json:"channel"`
	AuthorID  string `json:"author"`

	// Not a string because the API can return one of 10(!) objects *or* a string! Beautiful.
	// To get message content, assuming it's a string, use the m.String() method;
	// To get a MessageContent object, use m.Object().
	Content jsonutil.Raw `json:"content"`

	Attachments []Attachment `json:"attachments,omitempty"`
}

// String unmarshals m.Content to a string. Any errors are ignored.
func (m Message) String() string {
	var s string
	m.Content.UnmarshalTo(&s)
	return s
}

// Object unmarshals m.Content to a MessageContent.
func (m Message) Object() (c MessageContent, err error) {
	return c, m.Content.UnmarshalTo(&c)
}

// MessageContent is used for the Message.Content field, but only sometimes.
type MessageContent struct {
	Type string

	Content string `json:"content,omitempty"`
	Name    string `json:"name,omitempty"`
	// User ID
	ID string `json:"id,omitempty"`
	// User ID?
	By string `json:"by,omitempty"`
}

// AttachmentTag is used in the Attachment struct.
type AttachmentTag string

// Attachment tag constants. Not sure what these do?
const (
	TagAttachments AttachmentTag = "attachments"
	TagAvatars     AttachmentTag = "avatars"
	TagBackgrounds AttachmentTag = "backgrounds"
	TagBanners     AttachmentTag = "banners"
	TagIcons       AttachmentTag = "icons"
)

// Attachment is a message attachment.
type Attachment struct {
	ID  string        `json:"_id"`
	Tag AttachmentTag `json:"tag"`

	Size        uint64             `json:"size"` // Size in bytes
	Filename    string             `json:"filename"`
	Metadata    AttachmentMetadata `json:"metadata"`
	ContentType string             `json:"content_type"`
}

// AttachmentMetadata ...
type AttachmentMetadata struct {
	Type   string `json:"type"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}
