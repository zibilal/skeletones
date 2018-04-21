package order

import (
	"encoding/json"
	"fmt"
	"github.com/zibilal/skeletones/eventstore"
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
	EventBuilder eventstore.EventBuilder
}

func NewPlaceOrderCommand(order Order, eventBuilder eventstore.EventBuilder) *PlaceOrderCommand {
	placeOrder := new(PlaceOrderCommand)
	placeOrder.Status = "Placing Order"
	placeOrder.Timeplacement = time.Now()
	placeOrder.Order = order
	placeOrder.AggregateId = uuid.GenerateID()
	placeOrder.EventBuilder = eventBuilder
	return placeOrder
}

type OrderEventBuilder struct {
	ID          uuid.ID
	AggregateID uuid.ID
	Name        string
	Item        Order

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

	// Place order is create and event and aggregate
	evn, err := c.EventBuilder.SetName("order_placed").
		SetAggregateID(aggregateId).
			GenID().
				SetItem(ord).Build()

	if err != nil {
		return err
	}

	// put the event

	msgBus := eventstore.GetMessageBus()
	msgBus.Input() <- evn

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
	ProductId int
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
