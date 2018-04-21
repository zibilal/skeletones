package order

import (
	"github.com/zibilal/skeletones/uuid"
	"github.com/zibilal/skeletones/logger"
)

type OrderPlacedEvent struct {
	Name        string
	Order       Order
	ID          uuid.ID
	AggregateID uuid.ID
}

func NewOrderPlacedEvent(id uuid.ID, name string) *OrderPlacedEvent {
	o := new(OrderPlacedEvent)
	o.Name = name
	o.ID = id

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
	return nil
}
