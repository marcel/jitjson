package encoding

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/suite"
)

type BufferPoolTestSuite struct {
	suite.Suite
}

func TestBufferPoolTestSuite(t *testing.T) {
	suite.Run(t, new(BufferPoolTestSuite))
}

func (s *BufferPoolTestSuite) TestBufferPool() {
	bufSize := 256
	pool := NewSyncPool(bufSize)

	buf := pool.Get()
	s.Equal(bufSize, buf.Cap())

	buf.WriteString("1")

	pool.Put(buf)
	// Put resets the buffer
	buf1 := pool.Get()
	s.Equal(unsafe.Pointer(buf), unsafe.Pointer(buf1))
	s.Equal("", buf1.String())

	// We didn't put back buf1 so this is a different buffer
	buf2 := pool.Get()

	s.NotEqual(unsafe.Pointer(buf1), unsafe.Pointer(buf2))
}
