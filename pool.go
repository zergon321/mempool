package mempool

import "sync"

// Pool holds dynamically
// allocated objects that
// can be reused throughout
// the runtime.
type Pool[T Erasable] struct {
	objects     []T
	mut         *sync.Mutex
	constructor func() T
}

// Get extracts an empty object
// from the pool.
func (pool *Pool[T]) Get() Erasable {
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

// Put puts the object to the pool
// and erases its fields so it can
// be reused.
func (pool *Pool[T]) Put(object T) error {
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

// NewPool returns a new pool
// for a certain object type.
func NewPool[T Erasable](constructor func() T, options ...PoolOption[T]) (*Pool[T], error) {
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
	pool.constructor = constructor
	pool.objects = make([]T, 0, params.initCap)

	for i := 0; i < params.initLen; i++ {
		pool.objects = append(pool.objects, pool.constructor())
	}

	return &pool, nil
}
