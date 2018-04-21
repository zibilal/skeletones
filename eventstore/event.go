package eventstore

import (
	"fmt"
	"github.com/zibilal/skeletones/uuid"
	"sync"
)

// Event interface is base type for an event
type Event interface {
	Uuider
	fmt.Stringer
	Handle() error
}

type Uuider interface {
	SetID(id uuid.ID)
	GetID() uuid.ID
}

// Interface that decide how object will be encoded to
// and how object will be decoded from
// most likely object will be encoded to json bytes
// and  object will be decoded from json bytes
type Payloader interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}

// EventBuilder is interface for encapsulate the process
// of building an event
type EventBuilder interface {
	// Setup the item that an event carry
	// this is optional
	SetItem(interface{}) EventBuilder
	// Define the ID of an event
	// this is required
	// Build() function should return an error, if ID
	// has not initialized
	GenID() EventBuilder
	// AggregateID is generate the AggregateID of an event
	SetAggregateID(id uuid.ID) EventBuilder
	// Setup the name of the event
	// for maintenance reason, name of the event is being set
	// by the outside system
	SetName(string) EventBuilder
	// Build the event
	Build() (Event, error)
}

// EventCollection will save all accepted events for the app
// Event that will be saved in Event store should have been register
// in this event collection
type EventCollection struct {
	Collection map[string]Event
	Lock sync.Mutex
}

// NewEventCollection is factory function that generate
// the EventCollection object
func NewEventCollection() *EventCollection {
	collection := new(EventCollection)
	collection.Lock = sync.Mutex{}

	return collection
}

// RegisterEvent will assign name, with a particular event
func (c EventCollection) RegisterEvent(name string, event Event) {
	c.Lock.Lock()
	c.Collection[name] = event
	c.Lock.Unlock()
}

// GetEvent will fetch event with key name
func (c EventCollection) GetEvent(name string) (Event, error) {
	c.Lock.Lock()
	ev, found := c.Collection[name]
	c.Lock.Unlock()
	if !found {
		return nil, fmt.Errorf("could not find event of name %s", name)
	}

	return ev, nil
}

func (c EventCollection) PopEvent(name string) (Event, error) {
	c.Lock.Lock()
	ev, found := c.Collection[name]
	if found {
		delete(c.Collection, name)
	}
	c.Lock.Unlock()
	if !found {
		return nil, fmt.Errorf("could not find event of name %s", name)
	}

	return ev, nil
}