package mempool

// Erasable is an object
// whose data can be completely
// nullified in order to reuse it.
type Erasable interface {
	// Erase resets all the fields
	// of the object to defaults
	// (not deeply).
	Erase() error
}
