package gateway

// eventCreators is a map of event name : function to create a pointer
var eventCreators = map[string]func() interface{}{
	"Pong":          func() interface{} { return new(ping) },
	"Ready":         func() interface{} { return new(ReadyEvent) },
	"Message":       func() interface{} { return new(MessageEvent) },
	"MessageUpdate": func() interface{} { return new(MessageUpdateEvent) },
	"MessageDelete": func() interface{} { return new(MessageDeleteEvent) },
}
