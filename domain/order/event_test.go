package order

import (
	"github.com/zibilal/skeletones/persistence/inmemorypersistence"
	"github.com/zibilal/skeletones/uuid"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestNewOrderPlacedEvent(t *testing.T) {
	t.Log("Testing order placed event")
	{
		store := inmemorypersistence.NewInMemoryStore()
		orderPlacedEvent := NewOrderPlacedEvent(uuid.GenerateID(), "OrderPlacedEvent", store)

		if orderPlacedEvent.GetID() == "" {
			t.Errorf("%s expected ID not empty", failed)
		} else {
			vl := orderPlacedEvent.GetID().Valid()
			if vl {
				t.Logf("%s expected id is valid", success)
			} else {
				t.Errorf("%s expected id is valid got %b", failed, vl)
			}
		}

		if orderPlacedEvent.String() == "" {
			t.Errorf("%s expected Name not empty", failed)
		} else {
			t.Logf("%s expected Name not empty, got %s", success, orderPlacedEvent.String())
		}

		err := orderPlacedEvent.Handle()
		if err != nil {
			t.Errorf("%s expected error is nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error is nil", success)
		}

		dt, err := store.Fetch(orderPlacedEvent.AggregateID, nil)

		if err != nil {
			t.Errorf("%s expected error is nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error is nil", success)
		}

		if dt != nil {
			t.Logf("%s expected not nill, got %v", success, dt)
		} else {
			t.Errorf("%s expected not nil, got nil", failed)
		}
	}
}
