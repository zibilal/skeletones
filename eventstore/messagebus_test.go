package eventstore

import (
	"testing"
	"github.com/zibilal/skeletones/uuid"
	"fmt"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

type EventTest struct {
	ID uuid.ID
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
	fmt.Printf("\tHandling event %s, item %v", e.Name, e.Item)

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

		bus.Input() <- ev1
		bus.Input() <- ev2
		bus.Input() <- ev3

		t.Log("Ends")
	}
}
