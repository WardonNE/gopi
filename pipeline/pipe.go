package pipeline

type Handler[P, R any] func(passable P) R

type PipeHandler[P, R any] func(passable P, next Handler[P, R]) R

type IPipe[P, R any] interface {
	Handle(passable P, next Handler[P, R]) R
}

type pipe[P, R any] struct {
	handler PipeHandler[P, R]
}

func (s *pipe[P, R]) Handle(passable P, next Handler[P, R]) R {
	return s.handler(passable, next)
}

func AsPipe[P, R any](handler PipeHandler[P, R]) *pipe[P, R] {
	return &pipe[P, R]{
		handler: handler,
	}
}
