package eventbus

// ListenerClause is a listener callback
//
// example:
//
//  func OnLogin(event IEvent) bool {
//      // do something
//      return true
//  }
type ListenerClause = func(event EventInterface) bool

// Listener listens to one or more events
//
// non-static listener example:
//
//	type NonStaticListener struct {
//		Account string
//		Password string
//	}
//
//	func (l *NonStaticListener) New(data any) Listener {
//		values := data.(map[string]string)
//		return &NonStaticListener{
//		    Username: values["username"]
//	        Password: values["password"]
//		}
//	}
//
//	func (l *NonStaticListener) Handle(event IEvent) bool {
//		// do something
//		return true
//	}
//
// static listener example:
//
//  type StaticListener struct {
//
//  }
//
//  func (l *StaticListener) New(data any) Listener {
//      return l
//  }
//
//  func (l *StaticListener) Handle(event IEvent) bool {
//      // do something
//      return true
//  }
type Listener interface {
	// New creates a new [Listener] instance
	//
	// If listener instance is static, just return its self
	New(data any) Listener
	// Handle handles event and returns a bool value
	//
	// if false is returned, it will stop notify the followed listeners
	Handle(event EventInterface) bool
}

var _ Listener = (*clauseListener)(nil)

type clauseListener struct {
	clause *ListenerClause
}

func (l *clauseListener) New(data any) Listener {
	return l
}

func (l *clauseListener) Handle(event EventInterface) bool {
	return (*l.clause)(event)
}
