package encoding

import (
	"testing"

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
	s.Equal("1", pool.Get().String())

	buf2 := pool.Get()
	buf2.WriteString("2")

	s.NotEqual(buf.String, buf2.String())
}
