package eventbus

type Subscriber interface {
	Subscribe(eb IEventBus) map[string][]ListenerClause
}
