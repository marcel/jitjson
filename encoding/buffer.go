package encoding

import (
	"bytes"
	"strconv"
)

type Buffer struct {
	bytes.Buffer
	scratch    []byte
	scratchCap int
}

func NewBufferWithBuffer(b bytes.Buffer) *Buffer {
	buf := new(Buffer)
	buf.Buffer = b
	buf.scratchCap = 1024
	buf.scratch = make([]byte, buf.scratchCap)

	return buf
}

// TODO Unimplemented
// Invalid Kind = iota
// Uintptr
// Complex64
// Complex128
// Chan
// Func
// Interface
// Map
// Ptr
// Slice
// Struct
// UnsafePointer

func (e *Buffer) Bool(b bool) {
	e.WriteString(strconv.FormatBool(b))
}

func (e *Buffer) Float64(f float64) {
	encoded := strconv.AppendFloat(e.scratch[:0], f, 'E', -1, 64)
	e.Write(encoded)
}

func (e *Buffer) Float32(f float32) {
	e.Float64(float64(f))
}

func (e *Buffer) Int64(i int64) {
	encoded := strconv.AppendInt(e.scratch[:0], i, 10)
	e.Write(encoded)
}

func (e *Buffer) Int(i int) {
	e.Int64(int64(i))
}

func (e *Buffer) Int32(i int32) {
	e.Int64(int64(i))
}

func (e *Buffer) Int16(i int16) {
	e.Int64(int64(i))
}

func (e *Buffer) Int8(i int8) {
	e.Int64(int64(i))
}

func (e *Buffer) String(s string) {
	e.Quote(s)
}

func (e *Buffer) StringWithComma(s string) {
	offset := e.quote(s)
	e.scratch[offset] = ','
	e.Write(e.scratch[:offset+1])
}

func (e *Buffer) Uint64(ui uint64) {
	encoded := strconv.AppendUint(e.scratch[:0], ui, 10)
	e.Write(encoded)
}

func (e *Buffer) Uint32(ui uint32) {
	e.Uint64(uint64(ui))
}

func (e *Buffer) Uint16(ui uint16) {
	e.Uint64(uint64(ui))
}

func (e *Buffer) Uint8(ui uint8) {
	e.Uint64(uint64(ui))
}

func (e *Buffer) Uint(ui uint) {
	e.Uint64(uint64(ui))
}

// N.B. We aren't using WriteByte + WriteString + WriteByte or
// strconv.AppendQuote to avoid overhead of memory allocation and
// unecessary calls to bytes.Buffer.Grow
func (e *Buffer) Quote(s string) {
	strLen := len(s)
	if strLen > e.scratchCap {
		e.WriteByte('"')
		e.WriteString(s)
		e.WriteByte('"')
	} else {
		offset := e.quote(s)
		e.Write(e.scratch[:offset])
	}
}

// Writes quoted string to scratch byte slice returning offset but does
// not write it to the byte buffer yet so that callers can continue adding to
// the scratch byte slice before consolidating writes to the byte buffer
func (e *Buffer) quote(s string) int {
	strLen := len(s)
	offset := 0
	e.scratch[0] = '"'
	offset += strLen + 1
	copy(e.scratch[1:offset], s)
	e.scratch[offset] = '"'
	return offset + 1
}

func (e *Buffer) Comma() {
	e.WriteByte(',')
}

func (e *Buffer) Attr(name string) {
	offset := e.quote(name)
	e.scratch[offset] = ':'
	e.Write(e.scratch[:offset+1])
}

func (e *Buffer) OpenBrace() {
	e.WriteByte('{')
}

func (e *Buffer) CloseBrace() {
	e.WriteByte('}')
}
