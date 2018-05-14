package eventstore

import (
	"errors"
	"fmt"
	"github.com/zibilal/skeletones/logger"
	"github.com/zibilal/skeletones/uuid"
	"testing"
	"time"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

type EventTest struct {
	ID   uuid.ID
	Name string
	Item interface{}
}

func (e *EventTest) SetID(id uuid.ID) {
	e.ID = id
}

func (e *EventTest) GetID() uuid.ID {
	return e.ID
}

func (e *EventTest) String() string {
	return e.Name
}

func (e *EventTest) Handle() error {
	logger.Info(fmt.Sprintf("\tHandling event %s, item %v\n", e.Name, e.Item))

	if e.Name == "Ev3" {
		return errors.New("the event name cannot be Ev3")
	}

	return nil
}

func TestGetMessageBus(t *testing.T) {
	t.Log("Testing GetMessageBus")
	{
		bus := GetMessageBus()
		bus.SetBus(make(chan Event))

		ev1 := new(EventTest)
		ev1.Name = "Ev1"
		ev1.SetID(uuid.GenerateID())

		ev2 := new(EventTest)
		ev2.Name = "Ev2"
		ev2.SetID(uuid.GenerateID())

		ev3 := new(EventTest)
		ev3.Name = "Ev3"
		ev3.SetID(uuid.GenerateID())

		go func() {
			bus.Input() <- ev1
			bus.Input() <- ev2
			bus.Input() <- ev3
		}()

		go func() {
			bus.HandlingBus()
		}()

		time.Sleep(1 * time.Second)

		logger.Info("Starts")

		t.Log("Unit test ends")
	}
}
