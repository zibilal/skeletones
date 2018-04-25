package order

import (
	"github.com/zibilal/skeletones/logger"
	"github.com/zibilal/skeletones/persistence"
	"github.com/zibilal/skeletones/uuid"
)

type OrderPlacedEvent struct {
	Name        string
	Order       Order
	ID          uuid.ID
	AggregateID uuid.ID
	Store       persistence.Storer
}

func NewOrderPlacedEvent(id uuid.ID, name string, s persistence.Storer) *OrderPlacedEvent {
	o := new(OrderPlacedEvent)
	o.Name = name
	o.ID = id
	o.Store = s

	return o
}

func (e *OrderPlacedEvent) String() string {
	return e.Name
}

func (e *OrderPlacedEvent) SetID(id uuid.ID) {
	e.ID = id
}

func (e *OrderPlacedEvent) GetID() uuid.ID {
	return e.ID
}

func (e *OrderPlacedEvent) Handle() error {
	logger.Info("[Order Placed] Handle saving to write db")
	logger.Info("[Order Placed] Other things")

	// store the event
	return e.Store.Store(e.AggregateID, e)
}
