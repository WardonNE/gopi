package subscriber

import (
	"github.com/wardonne/gopi/eventbus"
)

type Subscriber interface {
	eventbus.Subscriber
	OnBeforeHandle(event eventbus.EventBus) bool
	OnAfterHandle(event eventbus.EventBus) bool
	OnFailedHandle(event eventbus.EventBus) bool
	OnRetryHandle(event eventbus.EventBus) bool
	OnProgressUpdated(event eventbus.EventBus) bool
}

type AbstractSubscriber struct {
}

func (subscriber *AbstractSubscriber) OnBeforeHandle(event eventbus.EventBus) bool {
	return true
}

func (subscriber *AbstractSubscriber) OnAfterHandle(event eventbus.EventBus) bool {
	return true
}

func (subscriber *AbstractSubscriber) OnFailedHandle(event eventbus.EventBus) bool {
	return true
}

func (subscriber *AbstractSubscriber) OnRetryHandle(event eventbus.EventBus) bool {
	return true
}

func (subscriber *AbstractSubscriber) OnProgressUpdated(event eventbus.EventBus) bool {
	return true
}
