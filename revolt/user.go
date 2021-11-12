package revolt

// User is a user.
type User struct {
	ID       string          `json:"_id"`
	Username string          `json:"username"`
	Avatar   *Attachment     `json:"avatar"`
	Bot      *BotInformation `json:"bot,omitempty"`
	Badges   Badges          `json:"badges"`
	Flags    UserFlags       `json:"flags"`

	Relations    []Relation `json:"relations"`
	Relationship string     `json:"relationship"`

	Status *Status `json:"status"`
	Online bool    `json:"online"`
}

// Badges ...
type Badges uint16

// Badges constants
const (
	BadgeDeveloper             Badges = 1 << 0
	BadgeTranslator            Badges = 1 << 1
	BadgeSupporter             Badges = 1 << 2
	BadgeResponsibleDisclosure Badges = 1 << 3
	BadgeRevoltTeam            Badges = 1 << 4

	BadgeEarlyAdopter Badges = 1 << 8
)

// UserFlags ...
type UserFlags uint16

// UserFlags constants
const (
	FlagSuspended UserFlags = 1 << 0
	FlagDeleted   UserFlags = 1 << 1
	FlagBanned    UserFlags = 1 << 2
)

// Relation ...
type Relation struct {
	User   string `json:"_id"`
	Status string `json:"status"`
}

// Status ...
type Status struct {
	Text     string `json:"text"`
	Presence string `json:"presence"`
}

// BotInformation ...
type BotInformation struct {
	OwnerID string `json:"owner"`
}
