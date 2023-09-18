package eventbus

var eventBus = NewEventBus()

// ListChannels lists all channels
func ListChannels() []string {
	return eventBus.ListChannels()
}

// GetChannel returns specific channel
func GetChannel(topic string) *EventChannel {
	return eventBus.GetChannel(topic)
}

// CreateChannel creates a new channel when it does not exists and returns the created channel
// If the specific channel already exists, it returns the channel
func CreateChannel(topic string) *EventChannel {
	return eventBus.CreateChannel(topic)
}

// DeleteChannel deletes specific channel
func DeleteChannel(topic string) {
	eventBus.DeleteChannel(topic)
}

// DispatchTo dispatches a new message to specific channel
func DispatchTo(topic string, data any) {
	eventBus.DispatchTo(topic, data)
}
