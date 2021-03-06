package order

import (
	"encoding/json"
	"fmt"
	"github.com/zibilal/skeletones/eventstore"
	"github.com/zibilal/skeletones/persistence"
	"github.com/zibilal/skeletones/uuid"
	"time"
)

type BaseCommand struct {
	Timeplacement time.Time
}

type PlaceOrderCommand struct {
	BaseCommand
	Status       string
	AggregateId  uuid.ID
	Order        Order
	Storer       persistence.Storer
	EventBuilder eventstore.EventBuilder
}

func NewPlaceOrderCommand(eventBuilder eventstore.EventBuilder, storer persistence.Storer) *PlaceOrderCommand {
	placeOrder := new(PlaceOrderCommand)
	placeOrder.Status = "Placing Order"
	placeOrder.Timeplacement = time.Now()
	placeOrder.AggregateId = uuid.GenerateID()
	placeOrder.EventBuilder = eventBuilder
	placeOrder.Storer = storer
	return placeOrder
}

type OrderEventBuilder struct {
	ID          uuid.ID
	AggregateID uuid.ID
	Name        string
	Item        Order
	Storer      persistence.Storer

	itemTmp interface{}
}

func NewOrderEventBuilder() *OrderEventBuilder {
	return new(OrderEventBuilder)
}

func (b *OrderEventBuilder) SetItem(itm interface{}) eventstore.EventBuilder {
	b.itemTmp = itm
	return b
}

func (b *OrderEventBuilder) GenID() eventstore.EventBuilder {
	b.ID = uuid.GenerateID()
	return b
}

func (b *OrderEventBuilder) SetAggregateID(id uuid.ID) eventstore.EventBuilder {
	if id == "" {
		b.AggregateID = uuid.GenerateID()
	} else {
		b.AggregateID = id
	}
	return b
}

func (b *OrderEventBuilder) SetName(name string) eventstore.EventBuilder {
	b.Name = name
	return b
}

func (b *OrderEventBuilder) SetStorer(storer persistence.Storer) eventstore.EventBuilder {
	b.Storer = storer
	return b
}

func (b *OrderEventBuilder) Build() (eventstore.Event, error) {
	if b.itemTmp != nil {
		switch b.itemTmp.(type) {
		case Order:
			b.Item = b.itemTmp.(Order)
		default:
			return nil, fmt.Errorf("unexpected item type %T, expected type Order", b.itemTmp)
		}
	}

	orderPlaced := &OrderPlacedEvent{
		ID:          b.ID,
		AggregateID: b.AggregateID,
		Name:        b.Name,
		Order:       b.Item,
		Store:       b.Storer,
	}

	return orderPlaced, nil
}

// command does not return anything
func (c *PlaceOrderCommand) Handle(payload eventstore.Payloader, aggregateId uuid.ID, revision int) error {
	// command create and save event in event store
	ord, found := payload.(Order)
	if !found {
		return fmt.Errorf("expected type Order, got %T", payload)
	}
	c.Order = ord

	// Place order is create and event and aggregate
	evn, err := c.EventBuilder.SetName("order_placed").
		SetAggregateID(aggregateId).
		GenID().
		SetItem(ord).
		SetStorer(c.Storer).Build()

	if err != nil {
		return err
	}

	// put the event

	go func() {
		msgBus := eventstore.GetMessageBus()
		msgBus.Input() <- evn
	}()

	return nil

}

type Order struct {
	OrderId int
	Status  string
	Items   []ProductVariant
}

// Encode the order type to json byte
func (o Order) Encode() ([]byte, error) {
	return json.Marshal(o)
}

// Decode the order type from json byte
func (o Order) Decode(b []byte) error {
	data := o
	return json.Unmarshal(b, &data)
}

type Variant struct {
	Color  string
	Size   int
	Width  int
	Height int
}

type ProductVariant struct {
	Sku       string
	Product   Product
	VariantId int
	Variant
	Price float64
}

type Product struct {
	ProductId   int
	Name        string
	Description string
	Category    string
}
