package eventbus

type Listener interface {
	Handle(data any) bool
}
