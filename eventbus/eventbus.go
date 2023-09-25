// eventbus is a package about handling events
package eventbus

import (
	"fmt"

	"github.com/wardonne/gopi/support/collection/list"
	"github.com/wardonne/gopi/support/maps"
)

// EventBusInterface is event bus interface
type EventBusInterface interface {
	// AddEvent registers an event to event bus.
	//
	// If topic exists, it will returns an error.
	AddEvent(event EventInterface) error
	// DeleteEvent unregisters an event from event bus.
	DeleteEvent(event EventInterface)
	// DeleteTopic unregisters an event from evnet bus by topic.
	DeleteTopic(topic string)
	// ListEvents lists all registered events.
	ListEvents() []EventInterface
	// ListTopics lists all topics.
	ListTopics() []string
	// OnEvent binds listener clauses to specific event.
	//
	// If the topic of event is not found, it will returns an error.
	OnEvent(event EventInterface, clauses []ListenerClause) error
	// OnTopic binds listener clauses to specific topic
	//
	// If the topic is not found, it will returns an error
	OnTopic(topic string, clauses []ListenerClause) error
	// SubscribeEvent binds listeners to specific event.
	//
	// If the topic of event is not found, it will returns an error.
	Listen(event EventInterface, listeners []Listener) error
	// SubscribeEvent binds listeners to specific topic.
	//
	// If topic is not found, it will returns an error
	ListenTopic(topic string, listeners []Listener) error
	// Subscribe adds an subscriber
	//
	// NOTICE: subscriber should be static
	Subscribe(subscriber Subscriber) error
	// Dispatch dispatches an event
	//
	// param `data` will pass into [Listener.New]
	Dispatch(event EventInterface, data any) error
}

var _ EventBusInterface = (*EventBus)(nil)

// EventBus is a basic implemention of [IEventBus]
type EventBus struct {
	events    *maps.SyncHashMap[string, EventInterface]
	listeners *maps.SyncHashMap[string, *list.ArrayList[Listener]]
}

// NewEventBus creates a new [EventBus] instance
func NewEventBus() *EventBus {
	return &EventBus{
		events: maps.NewSyncHashMap[string, EventInterface](),
	}
}

// ListEvents implements [IEventBus].ListEvents
func (eb *EventBus) ListEvents() []EventInterface {
	return eb.events.Values()
}

// ListTopics implements [IEventBus].ListTopics
func (eb *EventBus) ListTopics() []string {
	return eb.events.Keys()
}

// AddEvent implements [IEventBus].AddEvent
func (eb *EventBus) AddEvent(event EventInterface) error {
	if eb.events.ContainsKey(event.Topic()) {
		return fmt.Errorf("Topic \"%s\" exists", event.Topic())
	}
	eb.events.Set(event.Topic(), event)
	eb.listeners.Set(event.Topic(), list.NewArrayList[Listener]())
	return nil
}

// DeleteEvent implements [IEventBus].DeleteEvent
func (eb *EventBus) DeleteEvent(event EventInterface) {
	eb.events.Remove(event.Topic())
	eb.listeners.Remove(event.Topic())
}

// DeleteTopic implements [IEventBus].DeleteTopic
func (eb *EventBus) DeleteTopic(topic string) {
	eb.events.Remove(topic)
	eb.listeners.Remove(topic)
}

// OnEvent implements [IEventBus].OnEvent
func (eb *EventBus) OnEvent(event EventInterface, clauses []ListenerClause) error {
	if !eb.events.ContainsKey(event.Topic()) {
		return fmt.Errorf("Event not found")
	}
	var listenerList *list.ArrayList[Listener]
	if eb.listeners.ContainsKey(event.Topic()) {
		listenerList = eb.listeners.Get(event.Topic())
	} else {
		listenerList = list.NewArrayList[Listener]()
		eb.listeners.Set(event.Topic(), listenerList)
	}
	for _, clause := range clauses {
		listener := new(clauseListener)
		listener.clause = &clause
		listenerList.Add(listener)
	}
	return nil
}

// OnTopic implements [IEventBus].OnTopic
func (eb *EventBus) OnTopic(topic string, clauses []ListenerClause) error {
	if !eb.events.ContainsKey(topic) {
		return fmt.Errorf("Topic not found")
	}
	var listenerList *list.ArrayList[Listener]
	if eb.listeners.ContainsKey(topic) {
		listenerList = eb.listeners.Get(topic)
	} else {
		listenerList = list.NewArrayList[Listener]()
		eb.listeners.Set(topic, listenerList)
	}
	for _, clause := range clauses {
		listener := new(clauseListener)
		listener.clause = &clause
		listenerList.Add(listener)
	}
	return nil
}

// Listen implements [IEventBus].Listen
func (eb *EventBus) Listen(event EventInterface, listeners []Listener) error {
	if !eb.events.ContainsKey(event.Topic()) {
		return fmt.Errorf("Event not found")
	}
	if eb.listeners.ContainsKey(event.Topic()) {
		eb.listeners.Get(event.Topic()).AddAll(listeners...)
	} else {
		listenerList := list.NewArrayList[Listener]()
		listenerList.AddAll(listeners...)
		eb.listeners.Set(event.Topic(), listenerList)
	}
	return nil
}

// ListenTopic implements [IEventBus].ListenTopic
func (eb *EventBus) ListenTopic(topic string, listeners []Listener) error {
	if !eb.events.ContainsKey(topic) {
		return fmt.Errorf("Topic not found")
	}
	if eb.listeners.ContainsKey(topic) {
		eb.listeners.Get(topic).AddAll(listeners...)
	} else {
		listenerList := list.NewArrayList[Listener]()
		listenerList.AddAll(listeners...)
		eb.listeners.Set(topic, listenerList)
	}
	return nil
}

// Subscribe implements [IEventBus].Subscribe
func (eb *EventBus) Subscribe(subscriber Subscriber) error {
	listenerMap := subscriber.Subscribe(eb)
	for topic, listenerClauses := range listenerMap {
		if err := eb.OnTopic(topic, listenerClauses); err != nil {
			return err
		}
	}
	return nil
}

// Dispatch implements [IEventBus].Dispatch
func (eb *EventBus) Dispatch(event EventInterface, data any) error {
	if !eb.events.ContainsKey(event.Topic()) {
		return fmt.Errorf("Event not found")
	}
	listenerSet := eb.listeners.Get(event.Topic())
	listenerSet.Range(func(listener Listener) bool {
		return listener.New(data).Handle(event)
	})
	return nil
}
