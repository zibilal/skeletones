package eventstore

import (
	"fmt"
	"github.com/zibilal/skeletones/uuid"
)

type Event interface {
	Uuider
	fmt.Stringer
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
	ID() EventBuilder
	// Setup the name of the event
	// for maintenance reason, name of the event is being set
	// by the outside system
	SetName(string) EventBuilder
	// Build the event
	Build() (Event, error)
}
