package encoding

import (
	"bytes"
	"strconv"
)

type Buffer struct {
	bytes.Buffer
	scratch [64]byte
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
	if b {
		e.WriteString("true")
	} else {
		e.WriteString("false")
	}
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

func (e *Buffer) Quote(s string) {
	e.WriteByte('"')
	e.WriteString(s)
	e.WriteByte('"')
}

func (e *Buffer) Comma() {
	e.WriteByte(',')
}

func (e *Buffer) Attr(name string) {
	e.Quote(name)
	e.WriteByte(':')
}

func (e *Buffer) OpenBrace() {
	e.WriteByte('{')
}

func (e *Buffer) CloseBrace() {
	e.WriteByte('}')
}
