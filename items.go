package balabol

type ItemType string

const (
	TypeString = "string"
	TypeInt    = "uint64"
	TypeFloat  = "float"
	TypeJSON   = "json"
	TypeBLOB   = "blob"
)

type ItemValue interface {
	// Bytes converts item's value to a slice of bytes.
	Bytes() []byte
}

// Item is a high-level representation of a value in a storage.
type Item interface {
	// Value gets the item's value.
	Value() ItemValue
	// SetValue sets the item's value to the provided one.
	SetValue(value ItemValue) error
	// Path returns item's path
	Path() string
}

// item is a concrete implementation of Item interface.
type item struct {
	path  string
	value ItemValue
}

// Value returns item's value.
func (i *item) Value() ItemValue {
	return i.value
}

// SetValue allows changing of the item's value.
func (i *item) SetValue(value ItemValue) error {
	i.value = value
	return nil
}

// Path returns an item's path.
func (i *item) Path() string {
	return i.path
}

// NewItem is a factory method for Item.
func NewItem(path string, value ItemValue) Item {
	return &item{
		path:  path,
		value: value,
	}
}

// stringValue is a value of type String.
type stringValue struct {
	raw string
}

// NewStringValue returns new ItemValue that is a string.
func NewStringValue(value string) ItemValue {
	return &stringValue{raw: value}
}

// Bytes returns stringValue's contents as an array of bytes.
func (sv *stringValue) Bytes() []byte {
	return []byte(sv.raw)
}
