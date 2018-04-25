package inmemorypersistence

import (
	"github.com/zibilal/skeletones/uuid"
	"sync"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestNewInMemoryStore(t *testing.T) {
	t.Log("Testing InMemoryStore")
	{
		var wg sync.WaitGroup
		var dataInserted = make(chan bool)
		aggreagateId := uuid.GenerateID()

		inMemoryStore := NewInMemoryStore()

		wg.Add(1)
		go func() {
			defer wg.Done()

			<-dataInserted

			t.Log("Fetch data for aggregate id:", aggreagateId)

			data, err := inMemoryStore.Fetch(aggreagateId, nil)

			if err != nil {
				t.Errorf("%s expected error is nil, got %s", failed, err.Error())
			} else {
				t.Logf("%s expected error is nil", success)
			}

			if data == nil {
				t.Errorf("%s expected data is not nil for aggregate %s, got empty", failed, aggreagateId)
			} else {
				t.Logf("%s expected data is not nil for aggegate %s, got %v", success, aggreagateId, data)
			}
		}()

		wg.Add(1)
		go func() {

			defer wg.Done()

			dataExample1 := struct {
				ID      uuid.ID
				Name    string
				Address string
				Email   string
				Points  int
			}{
				ID:      uuid.GenerateID(),
				Name:    "Testing Name",
				Address: "Testing Address",
				Email:   "beexample@example.com",
				Points:  15,
			}

			dataExample2 := []struct {
				ID     uuid.ID
				Points int
			}{
				{
					ID:     "qwerqreqweqwerqwreqwe",
					Points: 17,
				},
				{
					ID:     "sadasdfasdfasdfsasdfas",
					Points: 18,
				},
			}

			storesData := []interface{}{
				dataExample1,
				dataExample2,
			}

			t.Log("[STORE]Storing for aggregate id:", aggreagateId)

			err := inMemoryStore.Store(aggreagateId, storesData)
			if err != nil {
				t.Errorf("%s [STORE]expected error is nil, got %s", failed, err.Error())
			} else {
				t.Logf("%s [STORE]expected error is nil", success)
			}

			dataInserted <- true
		}()

		wg.Wait()
	}
}
