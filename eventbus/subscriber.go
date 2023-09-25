package eventbus

// Subscriber is subscriber interface
type Subscriber interface {
	// Subscribe accepts an instance of IEventBus and returns listener map
	Subscribe(eb IEventBus) map[string][]ListenerClause
}
