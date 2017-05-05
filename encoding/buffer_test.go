package encoding

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type BufferTestSuite struct {
	suite.Suite
	buf Buffer
}

func TestBufferTestSuite(t *testing.T) {
	suite.Run(t, new(BufferTestSuite))
}

func (s *BufferTestSuite) SetupTest() {
	s.buf = *NewBufferWithBuffer(bytes.Buffer{})
}

func (s *BufferTestSuite) TestNewBufferWithBuffer() {
	underlying := bytes.Buffer{}
	underlying.Grow(256)

	buf := NewBufferWithBuffer(underlying)
	s.Equal(256, buf.Buffer.Cap())

	s.Equal(defaultScratchCap, buf.scratchCap)
	s.Equal(defaultScratchCap, cap(buf.scratch))
}

func (s *BufferTestSuite) TestBool() {
	s.buf.Bool(true)
	s.Equal("true", s.buf.String())

	s.buf.Reset()

	s.buf.Bool(false)
	s.Equal("false", s.buf.String())
}

func (s *BufferTestSuite) TestFloat() {
	var f64 float64 = 123456789.123456789
	s.buf.Float64(f64)
	s.Equal("1.2345678912345679E+08", s.buf.String())

	s.buf.Reset()

	f32 := float32(f64)
	s.buf.Float32(f32)
	s.Equal("1.23456792E+08", s.buf.String())
}

func (s *BufferTestSuite) TestInt() {
	s.buf.Int(100)
	s.buf.WriteByte(' ')
	s.buf.Int64(-64)
	s.buf.WriteByte(' ')
	s.buf.Int32(-32)
	s.buf.WriteByte(' ')
	s.buf.Int16(-16)
	s.buf.WriteByte(' ')
	s.buf.Int8(-8)
	s.buf.WriteByte(' ')
	s.buf.Uint(0)
	s.buf.WriteByte(' ')
	s.buf.Uint8(8)
	s.buf.WriteByte(' ')
	s.buf.Uint16(16)
	s.buf.WriteByte(' ')
	s.buf.Uint32(32)
	s.buf.WriteByte(' ')
	s.buf.Uint64(64)

	expected := "100 -64 -32 -16 -8 0 8 16 32 64"
	s.Equal(expected, s.buf.String())
}

func (s *BufferTestSuite) TestIntLargerThanCacheSize() {
	largeInt := intCacheSize + 1

	s.buf.Int64(largeInt)

	s.Equal(strconv.Itoa(int(largeInt)), s.buf.String())
}

func (s *BufferTestSuite) TestQuote() {
	str := "some text"
	s.buf.Quote(str)
	s.Equal(`"some text"`, s.buf.String())
}

func (s *BufferTestSuite) TestQuoteWithoutScratch() {
	buf := Buffer{}
	s.Equal(0, buf.scratchCap)

	buf.Quote("test")

	s.Equal(`"test"`, buf.String())
}

func (s *BufferTestSuite) TestComma() {
	s.buf.Comma()
	s.Equal(",", s.buf.String())
}

func (s *BufferTestSuite) TestBrace() {
	s.buf.OpenBrace()
	s.buf.CloseBrace()

	s.Equal("{}", s.buf.String())
}
