package command

import "github.com/zibilal/skeletones/eventstore"

// Basic interface for type that handle a command
type Handler interface {
	Handle(payloader eventstore.Payloader) error
}

// Command interface is the based typed for Command types
type Command interface {
	Handler
}
