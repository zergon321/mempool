package mempool

import "fmt"

type ErrorNegativeCapacity struct {
	capacity int
}

func (err *ErrorNegativeCapacity) Error() string {
	return fmt.Sprintf("got negative capacity: %d", err.capacity)
}

/*===============================================================*/

type ErrorNegativeLength struct {
	length int
}

func (err *ErrorNegativeLength) Error() string {
	return fmt.Sprintf("got negative length: %d", err.length)
}
