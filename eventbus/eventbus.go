package eventbus

import (
	"github.com/wardonne/gopi/support/collection/set"
	"github.com/wardonne/gopi/support/maps"
)

type EventBus struct {
	channels *maps.SyncHashMap[string, *EventChannel]
}

// NewEventBus creates a new event bus
func NewEventBus() *EventBus {
	return &EventBus{
		channels: maps.NewSyncHashMap[string, *EventChannel](),
	}
}

// GetChannel returns specific channel
func (eb *EventBus) ListChannels() []string {
	return eb.channels.Keys()
}

// GetChannel returns specific channel
func (eb *EventBus) GetChannel(topic string) *EventChannel {
	if eb.channels.ContainsKey(topic) {
		return eb.channels.Get(topic)
	}
	return nil
}

// CreateChannel creates a new channel when it does not exists and returns the created channel
// If the specific channel already exists, it returns the channel
func (eb *EventBus) CreateChannel(topic string) *EventChannel {
	if eb.channels.ContainsKey(topic) {
		return eb.channels.Get(topic)
	}
	channel := &EventChannel{
		listeners: set.NewHashSet[Listener](),
	}
	eb.channels.Set(topic, channel)
	return channel
}

// DeleteChannel deletes specific channel
func (eb *EventBus) DeleteChannel(topic string) {
	eb.channels.Remove(topic)
}

// DispatchTo dispatches a new message to specific channel
func (eb *EventBus) DispatchTo(topic string, data any) {
	if eb.channels.ContainsKey(topic) {
		channel := eb.channels.Get(topic)
		channel.Notify(data)
	}
}
