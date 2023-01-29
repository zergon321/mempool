package mempool

import "sync"

type PoolOption func(pool *Pool, params *poolParams) error

type poolParams struct {
	initCap int
	initLen int
}

func PoolOptionInitialCapacity(capacity int) PoolOption {
	return func(pool *Pool, params *poolParams) error {
		if capacity < 0 {
			return &ErrorNegativeCapacity{
				capacity: capacity,
			}
		}

		params.initCap = capacity
		return nil
	}
}

func PoolOptionInitialLength(length int) PoolOption {
	return func(pool *Pool, params *poolParams) error {
		if length < 0 {
			return &ErrorNegativeLength{
				length: length,
			}
		}

		params.initLen = length
		return nil
	}
}

func PoolOptionConcurrent() PoolOption {
	return func(pool *Pool, params *poolParams) error {
		pool.mut = &sync.Mutex{}
		return nil
	}
}
