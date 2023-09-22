package eventbus

import (
	"github.com/wardonne/gopi/support/maps"
)

type EventBus struct {
	events *maps.SyncHashMap[string, IEvent]
}

// NewEventBus creates a new event bus
func NewEventBus() *EventBus {
	return &EventBus{
		events: maps.NewSyncHashMap[string, IEvent](),
	}
}

// ListEvents lists all events
func (eb *EventBus) ListEvents() []string {
	return eb.events.Keys()
}

// GetEvent returns specific event
func (eb *EventBus) GetEvent(topic string) IEvent {
	if eb.events.ContainsKey(topic) {
		return eb.events.Get(topic)
	}
	return nil
}

// CreateEvent creates a new event when it does not exists and returns the created event
// If the specific event already exists, it returns the event
func (eb *EventBus) CreateEvent(event IEvent) IEvent {
	if eb.events.ContainsKey(event.Topic()) {
		return eb.events.Get(event.Topic())
	}
	eb.events.Set(event.Topic(), event)
	return event
}

// DeleteEvent deletes event by specific topic
func (eb *EventBus) DeleteEvent(topic string) {
	eb.events.Remove(topic)
}

// DispatchTo dispatches a new message to specific event by topic
func (eb *EventBus) DispatchTo(topic string, data any) {
	if eb.events.ContainsKey(topic) {
		event := eb.events.Get(topic)
		event.Notify(data)
	}
}
