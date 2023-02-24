package mempool

import "sync"

// PoolOption changes something
// in the newly created pool.
type PoolOption[T Erasable] func(pool *Pool[T], params *poolParams) error

// poolParams holds the
// parameters to be used
// for pool initialization.
type poolParams struct {
	initCap int // initCap is the initial capacity of the pool object slice.
	initLen int // initLen is the initial length the pool object slice.
}

// PoolOptionInitialCapacity creates an option
// to set the initial capacity of the object pool.
func PoolOptionInitialCapacity[T Erasable](capacity int) PoolOption[T] {
	return func(pool *Pool[T], params *poolParams) error {
		if capacity < 0 {
			return &ErrorNegativeCapacity{
				capacity: capacity,
			}
		}

		params.initCap = capacity
		return nil
	}
}

// PoolOptionInitialLength creates an option
// to set the initial length of the object pool.
func PoolOptionInitialLength[T Erasable](length int) PoolOption[T] {
	return func(pool *Pool[T], params *poolParams) error {
		if length < 0 {
			return &ErrorNegativeLength{
				length: length,
			}
		}

		params.initLen = length
		return nil
	}
}

// PoolOptionConcurrent enables a newly
// created pool to gandle concurrent gets and puts
// from many different goroutines at the same time.
func PoolOptionConcurrent[T Erasable]() PoolOption[T] {
	return func(pool *Pool[T], params *poolParams) error {
		pool.mut = &sync.Mutex{}
		return nil
	}
}
