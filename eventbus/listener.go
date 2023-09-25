package eventbus

type ListenerClause = func(event IEvent) bool

type Listener interface {
	// New creates a new [Listener] instance
	//
	// If listener instance is static, just return its self
	New(data any) Listener
	// Handle handles event and returns a bool value
	//
	// if false is returned, it will stop notify the followed listeners
	Handle(event IEvent) bool
}

var _ Listener = (*clauseListener)(nil)

type clauseListener struct {
	clause *ListenerClause
}

func (l *clauseListener) New(data any) Listener {
	return l
}

func (l *clauseListener) Handle(event IEvent) bool {
	return (*l.clause)(event)
}
