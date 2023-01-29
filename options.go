package mempool

import "sync"

type PoolOption[T Erasable] func(pool *Pool[T], params *poolParams) error

type poolParams struct {
	initCap int
	initLen int
}

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

func PoolOptionConcurrent[T Erasable]() PoolOption[T] {
	return func(pool *Pool[T], params *poolParams) error {
		pool.mut = &sync.Mutex{}
		return nil
	}
}
