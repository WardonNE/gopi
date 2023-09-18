package eventbus

import (
	"github.com/wardonne/gopi/support/collection/set"
)

type EventChannel struct {
	listeners *set.HashSet[Listener]
}

// Subscribe adds listeners to the channel.
func (channel *EventChannel) Subscribe(listeners ...Listener) *EventChannel {
	channel.listeners.AddAll(listeners...)
	return channel
}

// Notify notifies all the listeners.
func (channel *EventChannel) Notify(data any) {
	channel.listeners.Range(func(listener Listener) bool {
		return listener.Handle(data)
	})
}
