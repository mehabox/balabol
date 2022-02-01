package storage

import (
	"errors"
	"sync"

	"github.com/mehabox/balabol"
)

// memoryMapStorage is a proof-of-work storage engine.
// All it does is it saves and returns any items in application's memory.
type memoryMapStorage struct {
	sync.RWMutex
	db map[string]balabol.Item
}

func NewMemoryMapStorage() balabol.Storage {
	return &memoryMapStorage{db: make(map[string]balabol.Item)}
}

// Save saves an item to the map.
func (s *memoryMapStorage) Save(item balabol.Item) error {
	s.Lock()
	s.db[item.Path()] = item
	s.Unlock()
	return nil
}

// GetByPath gets an item by a path if it exists
func (s *memoryMapStorage) GetByPath(path string) (balabol.Item, error) {
	s.RLock()
	defer s.RUnlock()

	if item, ok := s.db[path]; ok {
		return item, nil
	}

	return nil, errors.New("couldn't find the item by path")
}

// Delete deletes an item from the storage.
func (s *memoryMapStorage) Delete(item balabol.Item) error {
	s.Lock()
	delete(s.db, item.Path())
	s.Unlock()

	return nil
}
