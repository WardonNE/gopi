package subscriber

import (
	"github.com/wardonne/gopi/eventbus"
	"github.com/wardonne/gopi/workerpool/event"
)

type SubscriberInterface interface {
	eventbus.Subscriber
	OnBeforeHandle(event eventbus.EventInterface) bool
	OnAfterHandle(event eventbus.EventInterface) bool
	OnFailedHandle(event eventbus.EventInterface) bool
	OnRetryHandle(event eventbus.EventInterface) bool
	OnProgressUpdated(event eventbus.EventInterface) bool
}

type Subscriber struct {
	BeforeHandle    func(eventbus.EventInterface) bool
	AfterHandle     func(eventbus.EventInterface) bool
	FailedHandle    func(eventbus.EventInterface) bool
	RetryHandle     func(eventbus.EventInterface) bool
	ProgressUpdated func(eventbus.EventInterface) bool
}

func (subscriber *Subscriber) OnBeforeHandle(event eventbus.EventInterface) bool {
	if subscriber.BeforeHandle != nil {
		return subscriber.BeforeHandle(event)
	}
	return true
}

func (subscriber *Subscriber) OnAfterHandle(event eventbus.EventInterface) bool {
	if subscriber.AfterHandle != nil {
		return subscriber.AfterHandle(event)
	}
	return true
}

func (subscriber *Subscriber) OnFailedHandle(event eventbus.EventInterface) bool {
	if subscriber.FailedHandle != nil {
		return subscriber.FailedHandle(event)
	}
	return true
}

func (subscriber *Subscriber) OnRetryHandle(event eventbus.EventInterface) bool {
	if subscriber.RetryHandle != nil {
		return subscriber.RetryHandle(event)
	}
	return true
}

func (subscriber *Subscriber) OnProgressUpdated(event eventbus.EventInterface) bool {
	if subscriber.ProgressUpdated != nil {
		return subscriber.ProgressUpdated(event)
	}
	return true
}

func (subscriber *Subscriber) Subscribe() map[string][]eventbus.ListenerClause {
	return map[string][]eventbus.ListenerClause{
		event.BeforeHandleTopic:    {subscriber.OnBeforeHandle},
		event.AfterHandleTopic:     {subscriber.OnAfterHandle},
		event.FailedHandleTopic:    {subscriber.OnFailedHandle},
		event.RetryHandleTopic:     {subscriber.OnRetryHandle},
		event.ProgressUpdatedTopic: {subscriber.OnProgressUpdated},
	}
}
