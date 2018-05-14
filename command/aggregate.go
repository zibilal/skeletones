package command

import "github.com/zibilal/skeletones/eventstore"

type Aggregator interface {
	ApplyEvents([]eventstore.Event) error
	ProcessCommand(command Command)
}
