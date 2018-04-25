package persistence

import (
	"github.com/zibilal/skeletones/uuid"
)

type Storer interface {
	Store(id uuid.ID, data interface{}) error
}

type Fetcher interface {
	Fetch(uuid.ID, func(interface{}) bool) (interface{}, error)
}
