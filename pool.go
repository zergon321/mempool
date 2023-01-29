package mempool

import "sync"

type Pool[T Erasable] struct {
	objects []*T
	mut     *sync.Mutex
}

func (pool *Pool[T]) Get() *T {
	if len(pool.objects) <= 0 {
		var zeroVal T
		return &zeroVal
	}

	if pool.mut != nil {
		pool.mut.Lock()
		defer pool.mut.Unlock()
	}

	object := pool.objects[len(pool.objects)-1]
	pool.objects = pool.objects[:len(pool.objects)-1]

	return object
}

func (pool *Pool[T]) Put(object *T) error {
	if pool.mut != nil {
		pool.mut.Lock()
		defer pool.mut.Unlock()
	}

	err := (*object).Erase()

	if err != nil {
		return err
	}

	pool.objects = append(pool.objects, object)

	return nil
}

func NewPool[T Erasable](options ...PoolOption[T]) (*Pool[T], error) {
	var pool Pool[T]
	var params poolParams

	// Apply all the options.
	for i := 0; i < len(options); i++ {
		option := options[i]
		err := option(&pool, &params)

		if err != nil {
			return nil, err
		}
	}

	// Fill the pool.
	pool.objects = make([]*T, 0, params.initCap)

	for i := 0; i < params.initLen; i++ {
		var zeroVal T
		pool.objects = append(pool.objects, &zeroVal)
	}

	return &pool, nil
}
