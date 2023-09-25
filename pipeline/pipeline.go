package pipeline

type Pipeline[P, R any] struct {
	passable    P
	pipes       []IPipe[P, R]
	destination func(passable P) R
}

func NewPipeline[P, R any]() *Pipeline[P, R] {
	return &Pipeline[P, R]{}
}

func (p *Pipeline[P, R]) Send(passable P) *Pipeline[P, R] {
	p.passable = passable
	return p
}

func (p *Pipeline[P, R]) Through(pipes ...IPipe[P, R]) *Pipeline[P, R] {
	p.pipes = pipes
	return p
}

func (p *Pipeline[P, R]) AppendThrough(pipe IPipe[P, R]) *Pipeline[P, R] {
	p.pipes = append(p.pipes, pipe)
	return p
}

func (p *Pipeline[P, R]) Then(destination Callback[P, R]) R {
	p.destination = destination
	stack := p.destination
	length := len(p.pipes)
	for i := length; i > 0; i-- {
		pipe := p.pipes[i-1]
		stack = func(s func(passable P) R, pipe IPipe[P, R]) func(passable P) R {
			return func(passable P) R {
				return pipe.Handle(passable, s)
			}
		}(stack, pipe)
	}
	return stack(p.passable)
}
