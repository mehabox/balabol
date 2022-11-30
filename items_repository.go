package balabol

// ItemsRepository is a repository of Items.
type ItemsRepository interface {
	// GetByPath returns an item by its path.
	GetByPath(path string) (Item, error)
	// Save saves an item.
	Save(item Item) error
	// Delete deletes an item.
	Delete(item Item) error
}
