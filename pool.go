package mempool

import (
	"sync"

	ll "github.com/zergon321/ll"
)

// Pool holds dynamically
// allocated objects that
// can be reused throughout
// the runtime.
type Pool[T Erasable] struct {
	objects     *ll.AmortizedList[T]
	mut         *sync.Mutex
	constructor func() T
}

// Get extracts an empty object
// from the pool.
func (pool *Pool[T]) Get() T {
	if pool.objects.Len() <= 0 {
		return pool.constructor()
	}

	if pool.mut != nil {
		pool.mut.Lock()
		defer pool.mut.Unlock()
	}

	object := pool.objects.Remove(pool.objects.Back())

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

	pool.objects.PushBack(object)

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
	pool.objects = ll.NewAmortized[T]()

	for i := 0; i < params.initLen; i++ {
		pool.objects.PushBack(pool.constructor())
	}

	return &pool, nil
}
