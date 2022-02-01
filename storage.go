package balabol

// Storage depicts a lower-level storage engine/driver used to hold and organize Items.
type Storage interface {
	// Save saves an item to the storage.
	Save(item Item) error
	// GetByPath returns an Item if it exists on the path, otherwise, an error.
	GetByPath(path string) (Item, error)
	// Delete permanently deletes an Item from the storage.
	Delete(item Item) error
}
