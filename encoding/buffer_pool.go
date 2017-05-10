package encoding

import (
	"bytes"
	"sync"
)

// BufferPool provides an API for managing
// bytes.Buffer objects:
type BufferPool interface {
	Get() *Buffer
	Put(*Buffer)
}

// syncPoolBufPool is an implementation of BufferPool
// that uses a sync.Pool to maintain buffers:
type syncPoolBufPool struct {
	pool       *sync.Pool
	makeBuffer func() interface{}
}

func NewSyncPool(bufSize int) BufferPool {
	var newPool syncPoolBufPool

	newPool.makeBuffer = func() interface{} {
		var b bytes.Buffer
		b.Grow(bufSize)

		return NewBufferWithBuffer(b)
	}

	newPool.pool = &sync.Pool{}
	newPool.pool.New = newPool.makeBuffer

	return &newPool
}

func (bp *syncPoolBufPool) Get() (b *Buffer) {
	poolObject := bp.pool.Get()

	b, ok := poolObject.(*Buffer)
	if !ok {
		return bp.makeBuffer().(*Buffer)
	}
	return b
}

func (bp *syncPoolBufPool) Put(b *Buffer) {
	b.Reset()
	bp.pool.Put(b)
}
