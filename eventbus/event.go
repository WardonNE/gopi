package eventbus

// EventInterface is the interface of event
type EventInterface interface {
	// Topic returns the unique topic
	Topic() string
}
