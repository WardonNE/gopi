package pipeline

type Callback[P, R any] func(passable P) R

type Handler[P, R any] func(passable P, next Callback[P, R]) R

type IPipe[P, R any] interface {
	Handle(passable P, next Callback[P, R]) R
}

type pipe[P, R any] struct {
	handler Handler[P, R]
}

func (s *pipe[P, R]) Handle(passable P, next Callback[P, R]) R {
	return s.handler(passable, next)
}

func AsPipe[P, R any](handler Handler[P, R]) *pipe[P, R] {
	return &pipe[P, R]{
		handler: handler,
	}
}
