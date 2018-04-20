package order

import (
	"encoding/json"
	"github.com/zibilal/skeletones/uuid"
	"time"
)

type BaseCommand struct {
	Timeplacement time.Time
}

type PlaceOrderCommand struct {
	BaseCommand
	Status      string
	AggregateId uuid.ID
	Order       Order
}

func NewPlaceOrderCommand(order Order) *PlaceOrderCommand {
	placeOrder := new(PlaceOrderCommand)
	placeOrder.Status = "Placing Order"
	placeOrder.Timeplacement = time.Now()
	placeOrder.Order = order
	placeOrder.AggregateId = uuid.GenerateID()
	return placeOrder
}

// command does not return anything
func (c *PlaceOrderCommand) Handle() {
	// should command create and save event in event store
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
