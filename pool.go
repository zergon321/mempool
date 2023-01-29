package mempool

import "sync"

type Pool struct {
	objects     []Erasable
	mut         *sync.Mutex
	constructor func() Erasable
}

func (pool *Pool) Get() Erasable {
	if len(pool.objects) <= 0 {
		return pool.constructor()
	}

	if pool.mut != nil {
		pool.mut.Lock()
		defer pool.mut.Unlock()
	}

	object := pool.objects[len(pool.objects)-1]
	pool.objects = pool.objects[:len(pool.objects)-1]

	return object
}

func (pool *Pool) Put(object Erasable) error {
	if pool.mut != nil {
		pool.mut.Lock()
		defer pool.mut.Unlock()
	}

	err := object.Erase()

	if err != nil {
		return err
	}

	pool.objects = append(pool.objects, object)

	return nil
}

func NewPool[T Erasable](constructor func() Erasable, options ...PoolOption) (*Pool, error) {
	var pool Pool
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
	pool.constructor = constructor
	pool.objects = make([]Erasable, 0, params.initCap)

	for i := 0; i < params.initLen; i++ {
		pool.objects = append(pool.objects, pool.constructor())
	}

	return &pool, nil
}
