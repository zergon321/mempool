package mempool

import "fmt"

// ErrorNegativeCapacity is returned if
// a negative value has been passed for
// pool capacity as an option.
type ErrorNegativeCapacity struct {
	capacity int
}

// Error returns the error message.
func (err *ErrorNegativeCapacity) Error() string {
	return fmt.Sprintf("got negative capacity: %d", err.capacity)
}

/*===============================================================*/

// ErrorNegativeLength is returned if
// a negative value has been passed for
// pool length as an option.
type ErrorNegativeLength struct {
	length int
}

// Error returns the error message.
func (err *ErrorNegativeLength) Error() string {
	return fmt.Sprintf("got negative length: %d", err.length)
}
