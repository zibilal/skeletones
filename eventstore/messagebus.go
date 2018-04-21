package eventstore

import (
	"github.com/zibilal/skeletones/logger"
	"sync"
)

type MessageBus struct {
	EventBus chan Event
	CanRun   bool
}

var bus *MessageBus
var once sync.Once

func GetMessageBus() *MessageBus {
	once.Do(func() {
		bus = new(MessageBus)
		bus.CanRun = true
	})
	return bus
}

func (b *MessageBus) SetBus(ch chan Event) {
	b.EventBus = ch
}

func (b *MessageBus) Input() chan<- Event {
	return b.EventBus
}

func (b *MessageBus) HandlingBus() {
	go func() {
		for b.CanRun {
			evn := <-b.EventBus
			if err := evn.Handle(); err != nil {
				logger.Info("failed handling", evn.String(), "caused", err.Error())
				return
			}
		}
	}()
}
