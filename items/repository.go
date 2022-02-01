package items

import "github.com/mehabox/balabol"

type itemsRepository struct {
	storage balabol.Storage
}

// GetByPath returns an item by its path
func (r *itemsRepository) GetByPath(path string) (balabol.Item, error) {
	return r.storage.GetByPath(path)
}

// Save saves an item to the storage.
func (r *itemsRepository) Save(item balabol.Item) error {
	return r.storage.Save(item)
}

// Delete deletes an item from the storage engine.
func (r *itemsRepository) Delete(item balabol.Item) error {
	return r.storage.Delete(item)
}

// NewItemsRepository initializes a new repository of items.
func NewItemsRepository(storage balabol.Storage) balabol.ItemsRepository {
	return &itemsRepository{storage: storage}
}
