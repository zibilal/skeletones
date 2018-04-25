package inmemorypersistence

import (
	"github.com/zibilal/skeletones/uuid"
	"sync"
	"fmt"
)

type InMemoryStore struct {
	data     map[uuid.ID]interface{}
	dataLock sync.Mutex
}

func NewInMemoryStore() *InMemoryStore {
	inMemory := new(InMemoryStore)
	inMemory.data = make(map[uuid.ID]interface{})
	return inMemory
}

func (m *InMemoryStore) Store(aggregateId uuid.ID, data interface{}) error {
	m.dataLock.Lock()
	m.data[aggregateId] = data
	m.dataLock.Unlock()

	return nil
}

func (m *InMemoryStore) Fetch(aggregateId uuid.ID, filter func(interface{}) bool) (interface{}, error) {
	m.dataLock.Lock()
	dt, found := m.data[aggregateId]
	m.dataLock.Unlock()

	if !found {
		return nil, fmt.Errorf("data with aggregate id %s not found", aggregateId)
	}

	return dt, nil
}
