package eventbus

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testevent1 struct {
	Value string
}

func (t *testevent1) Topic() string {
	return "testevent1"
}

type testevent1listener struct{}

func (l *testevent1listener) New(data any) Listener { return l }

func (l *testevent1listener) Handle(event EventInterface) bool {
	event.(*testevent1).Value += "testevent1listener"
	return true
}

type testevent2 struct {
	Value string
}

func (t *testevent2) Topic() string {
	return "testevent2"
}

type testevent2listener struct{ Value string }

func (l *testevent2listener) New(data any) Listener { return &testevent2listener{Value: data.(string)} }

func (l *testevent2listener) Handle(event EventInterface) bool {
	event.(*testevent2).Value += l.Value
	return true
}

type unregisteredevent struct{}

func (t *unregisteredevent) Topic() string {
	return "unregisteredevent"
}

type testsubscriber struct{}

func (t *testsubscriber) testevent1(event EventInterface) bool {
	event.(*testevent1).Value += "testsubscriber"
	return true
}

func (t *testsubscriber) testevent2(event EventInterface) bool {
	event.(*testevent2).Value += "testsubscriber"
	return true
}

func (t *testsubscriber) Subscribe() map[string][]ListenerClause {
	return map[string][]func(event EventInterface) bool{
		"testevent1": []ListenerClause{t.testevent1},
		"testevent2": []ListenerClause{t.testevent2},
	}
}

type errtestsubscriber struct{}

func (t *errtestsubscriber) Subscribe() map[string][]ListenerClause {
	return map[string][]func(event EventInterface) bool{
		"unregisteredevent": []ListenerClause{},
	}
}

func TestEventBus(t *testing.T) {
	eb := NewEventBus()
	assert.Nil(t, eb.AddEvent(&testevent1{}))
	assert.Nil(t, eb.AddEvent(&testevent2{}))
	assert.Equal(t, fmt.Errorf("Topic \"%s\" exists", (&testevent1{}).Topic()).Error(), eb.AddEvent(&testevent1{}).Error())
	assert.ElementsMatch(t, []EventInterface{&testevent1{}, &testevent2{}}, eb.ListEvents())
	assert.ElementsMatch(t, []string{"testevent1", "testevent2"}, eb.ListTopics())

	eb.DeleteEvent(&testevent2{})
	assert.ElementsMatch(t, []EventInterface{&testevent1{}}, eb.ListEvents())
	assert.ElementsMatch(t, []string{"testevent1"}, eb.ListTopics())

	assert.Nil(t, eb.AddEvent(&testevent2{}))

	eb.DeleteTopic("testevent2")
	assert.ElementsMatch(t, []EventInterface{&testevent1{}}, eb.ListEvents())
	assert.ElementsMatch(t, []string{"testevent1"}, eb.ListTopics())

	assert.Nil(t, eb.AddEvent(&testevent2{}))

	assert.Nil(t, eb.OnEvent(&testevent1{}, []ListenerClause{
		func(event EventInterface) bool {
			event.(*testevent1).Value += "listenerclause"
			return true
		},
	}))
	assert.Nil(t, eb.OnTopic("testevent2", []ListenerClause{
		func(event EventInterface) bool {
			event.(*testevent2).Value += "listenerclause"
			return true
		},
	}))
	assert.Equal(t, "Event not found", eb.OnEvent(&unregisteredevent{}, []ListenerClause{func(event EventInterface) bool {
		return true
	}}).Error())
	assert.Equal(t, "Topic not found", eb.OnTopic("unregisteredevent", []ListenerClause{func(event EventInterface) bool { return true }}).Error())

	assert.Nil(t, eb.Listen(&testevent1{}, []Listener{&testevent1listener{}}))
	assert.Nil(t, eb.ListenTopic("testevent2", []Listener{&testevent2listener{}}))
	assert.Equal(t, "Event not found", eb.Listen(&unregisteredevent{}, []Listener{&testevent1listener{}}).Error())
	assert.Equal(t, "Topic not found", eb.ListenTopic("unregisteredevent", []Listener{&testevent1listener{}}).Error())

	assert.Nil(t, eb.Subscribe(&testsubscriber{}))
	assert.Equal(t, "Topic not found", eb.Subscribe(&errtestsubscriber{}).Error())

	evt1 := &testevent1{Value: "testevent1"}
	assert.Nil(t, eb.Dispatch(evt1, nil))
	assert.Equal(t, "testevent1listenerclausetestevent1listenertestsubscriber", evt1.Value)

	evt2 := &testevent2{Value: "testevent2"}
	assert.Nil(t, eb.Dispatch(evt2, "initdata"))
	assert.Equal(t, "testevent2listenerclauseinitdatatestsubscriber", evt2.Value)

	assert.Equal(t, "Event not found", eb.Dispatch(&unregisteredevent{}, nil).Error())
}
