package eventbus

// IEvent is the interface of event
type IEvent interface {
	// Topic returns the unique topic
	Topic() string
}
