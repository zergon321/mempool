package mempool_test

import (
	"sync"
	"testing"

	"github.com/zergon321/mempool"
)

type Data struct {
	X float32
	Y float32
}

func (data *Data) Erase() error {
	data.X = 0
	data.Y = 0

	return nil
}

func BenchmarkSyncPool(b *testing.B) {
	pool := &sync.Pool{
		New: func() any { return &Data{X: 32, Y: 16} },
	}

	for i := 0; i < b.N; i++ {
		pool.Put(&Data{X: 32, Y: 16})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkSyncPoolFill(b *testing.B) {
	pool := &sync.Pool{
		New: func() any { return &Data{X: 32, Y: 16} },
	}

	for i := 0; i < b.N; i++ {
		pool.Put(&Data{X: 32, Y: 16})
	}
}

func BenchmarkSyncPoolRefill(b *testing.B) {
	pool := &sync.Pool{
		New: func() any { return &Data{X: 32, Y: 16} },
	}

	for i := 0; i < b.N; i++ {
		pool.Put(&Data{X: 32, Y: 16})
	}

	data := make([]*Data, 0, b.N)

	for i := 0; i < b.N; i++ {
		data = append(data, pool.Get().(*Data))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.Put(data[i])
	}
}

func BenchmarkMempool(b *testing.B) {
	pool, _ := mempool.NewPool[*Data](
		func() *Data { return &Data{X: 32, Y: 16} })

	for i := 0; i < b.N; i++ {
		pool.Put(&Data{X: 32, Y: 16})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkMempoolFill(b *testing.B) {
	pool, _ := mempool.NewPool[*Data](
		func() *Data { return &Data{X: 32, Y: 16} })

	for i := 0; i < b.N; i++ {
		pool.Put(&Data{X: 32, Y: 16})
	}
}

func BenchmarkMempoolRefill(b *testing.B) {
	pool, _ := mempool.NewPool[*Data](
		func() *Data { return &Data{X: 32, Y: 16} })

	for i := 0; i < b.N; i++ {
		pool.Put(&Data{X: 32, Y: 16})
	}

	data := make([]*Data, 0, b.N)

	for i := 0; i < b.N; i++ {
		data = append(data, pool.Get())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.Put(data[i])
	}
}

func BenchmarkMempoolConcurrent(b *testing.B) {
	pool, _ := mempool.NewPool[*Data](
		func() *Data { return &Data{X: 32, Y: 16} },
		mempool.PoolOptionConcurrent[*Data]())

	for i := 0; i < b.N; i++ {
		pool.Put(&Data{X: 32, Y: 16})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkMempoolFillConcurrent(b *testing.B) {
	pool, _ := mempool.NewPool[*Data](
		func() *Data { return &Data{X: 32, Y: 16} },
		mempool.PoolOptionConcurrent[*Data]())

	for i := 0; i < b.N; i++ {
		pool.Put(&Data{X: 32, Y: 16})
	}
}

func BenchmarkMempoolRefillConcurrent(b *testing.B) {
	pool, _ := mempool.NewPool[*Data](
		func() *Data { return &Data{X: 32, Y: 16} },
		mempool.PoolOptionConcurrent[*Data]())

	for i := 0; i < b.N; i++ {
		pool.Put(&Data{X: 32, Y: 16})
	}

	data := make([]*Data, 0, b.N)

	for i := 0; i < b.N; i++ {
		data = append(data, pool.Get())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.Put(data[i])
	}
}
