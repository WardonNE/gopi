package subscriber

import (
	"github.com/wardonne/gopi/eventbus"
	"github.com/wardonne/gopi/workerpool/event"
)

// Interface subscriber interface
type Interface interface {
	eventbus.Subscriber
	OnBeforeHandle(event eventbus.EventInterface) bool
	OnAfterHandle(event eventbus.EventInterface) bool
	OnFailedHandle(event eventbus.EventInterface) bool
	OnRetryHandle(event eventbus.EventInterface) bool
}

// Subscriber subscriber
type Subscriber struct {
	BeforeHandle    func(eventbus.EventInterface) bool
	AfterHandle     func(eventbus.EventInterface) bool
	FailedHandle    func(eventbus.EventInterface) bool
	RetryHandle     func(eventbus.EventInterface) bool
	ProgressUpdated func(eventbus.EventInterface) bool
}

// OnBeforeHandle handles on before event
func (subscriber *Subscriber) OnBeforeHandle(event eventbus.EventInterface) bool {
	if subscriber.BeforeHandle != nil {
		return subscriber.BeforeHandle(event)
	}
	return true
}

// OnAfterHandle handles on after event
func (subscriber *Subscriber) OnAfterHandle(event eventbus.EventInterface) bool {
	if subscriber.AfterHandle != nil {
		return subscriber.AfterHandle(event)
	}
	return true
}

// OnFailedHandle handles on failed event
func (subscriber *Subscriber) OnFailedHandle(event eventbus.EventInterface) bool {
	if subscriber.FailedHandle != nil {
		return subscriber.FailedHandle(event)
	}
	return true
}

// OnRetryHandle handles on retry event
func (subscriber *Subscriber) OnRetryHandle(event eventbus.EventInterface) bool {
	if subscriber.RetryHandle != nil {
		return subscriber.RetryHandle(event)
	}
	return true
}

// Subscribe returns top-event map
func (subscriber *Subscriber) Subscribe() map[string][]eventbus.ListenerClause {
	return map[string][]eventbus.ListenerClause{
		event.BeforeHandleTopic: {subscriber.OnBeforeHandle},
		event.AfterHandleTopic:  {subscriber.OnAfterHandle},
		event.FailedHandleTopic: {subscriber.OnFailedHandle},
		event.RetryHandleTopic:  {subscriber.OnRetryHandle},
	}
}
