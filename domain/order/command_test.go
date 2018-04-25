package order

import (
	"testing"
	"github.com/zibilal/skeletones/uuid"
	"github.com/zibilal/skeletones/eventstore"
	"time"
	"github.com/zibilal/skeletones/logger"
	"github.com/zibilal/skeletones/persistence/inmemorypersistence"
)

func TestNewPlaceOrderCommand(t *testing.T) {
	t.Log("Testing NewPlaceOrderCommand")
	{
		orderTest := Order{
			OrderId: 1,
			Status: "OrderPlaced",
			Items: []ProductVariant {

			},
		}

		orderEventBuilder := NewOrderEventBuilder()
		cmd := NewPlaceOrderCommand(orderEventBuilder, inmemorypersistence.NewInMemoryStore())
		err := cmd.Handle(orderTest, uuid.GenerateID(), 0)
		if err != nil {
			t.Errorf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		msgBus := eventstore.GetMessageBus()
		msgBus.SetBus(make(chan eventstore.Event))
		msgBus.HandlingBus()

		time.Sleep(1 * time.Second)

		logger.Info("--")

		t.Log("Command test ends")
	}
}
