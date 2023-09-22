package eventbus

var eventBus = NewEventBus()

// ListEvents lists all events
func ListEvents() []string {
	return eventBus.ListEvents()
}

// GetEvent returns specific event
func GetEvent(topic string) IEvent {
	return eventBus.GetEvent(topic)
}

// CreateEvent creates a new event when it does not exists and returns the created event
// If the specific event already exists, it returns the event
func CreateEvent(event IEvent) IEvent {
	return eventBus.CreateEvent(event)
}

// DeleteEvent deletes specific event
func DeleteEvent(topic string) {
	eventBus.DeleteEvent(topic)
}

// DispatchTo dispatches a new message to specific event
func DispatchTo(topic string, data any) {
	eventBus.DispatchTo(topic, data)
}
