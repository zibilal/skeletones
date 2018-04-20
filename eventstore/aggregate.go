package eventstore

import "github.com/zibilal/skeletones/command"

type Aggregator interface {
	ApplyEvents([]Event) error
	ProcessCommand(command command.Command)
}
