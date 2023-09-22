package eventbus

// IEvent is the interface of event
type IEvent interface {
	// Topic returns the unique topic
	Topic() string
	// AddListener add listeners to the event
	AddListener(listeners ...Listener)
	// Notify notifies all the listeners
	Notify(data any)
}
