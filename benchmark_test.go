package mempool_test

import (
	"sync"
	"testing"

	"github.com/zergon321/mempool"
)

type ByteArray []byte

func (array ByteArray) Erase() error {
	for i := 0; i < len(array); i++ {
		array[i] = 0
	}

	return nil
}

func BenchmarkSyncPool(b *testing.B) {
	pool := &sync.Pool{
		New: func() any { return make([]byte, 64) },
	}

	for i := 0; i < b.N; i++ {
		pool.Put(make([]byte, 64))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkSyncPoolFill(b *testing.B) {
	pool := &sync.Pool{
		New: func() any { return make([]byte, 64) },
	}

	for i := 0; i < b.N; i++ {
		pool.Put(make([]byte, 64))
	}
}

func BenchmarkMempool(b *testing.B) {
	pool, _ := mempool.NewPool[ByteArray](
		func() ByteArray { return make([]byte, 64) })

	for i := 0; i < b.N; i++ {
		pool.Put(make([]byte, 64))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkMempoolFill(b *testing.B) {
	pool, _ := mempool.NewPool[ByteArray](
		func() ByteArray { return make([]byte, 64) })

	for i := 0; i < b.N; i++ {
		pool.Put(make([]byte, 64))
	}
}

func BenchmarkMempoolRefill(b *testing.B) {
	pool, _ := mempool.NewPool[ByteArray](
		func() ByteArray { return make([]byte, 64) })

	for i := 0; i < b.N; i++ {
		pool.Put(make([]byte, 64))
	}

	data := make([]ByteArray, 0, b.N)

	for i := 0; i < b.N; i++ {
		data = append(data, pool.Get())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.Put(data[i])
	}
}
