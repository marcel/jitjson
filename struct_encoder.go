package jitjson

import (
	"bytes"
	"strconv"
)

type structEncoder struct {
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

func (e *structEncoder) bool(b bool) {
	if b {
		e.WriteString("true")
	} else {
		e.WriteString("false")
	}
}

func (e *structEncoder) float64(f float64) {
	encoded := strconv.AppendFloat(e.scratch[:0], f, 'E', -1, 64)
	e.Write(encoded)
}

func (e *structEncoder) float32(f float32) {
	e.float64(float64(f))
}

func (e *structEncoder) int64(i int64) {
	encoded := strconv.AppendInt(e.scratch[:0], i, 10)
	e.Write(encoded)
}

func (e *structEncoder) int(i int) {
	e.int64(int64(i))
}

func (e *structEncoder) int32(i int32) {
	e.int64(int64(i))
}

func (e *structEncoder) int16(i int16) {
	e.int64(int64(i))
}

func (e *structEncoder) int8(i int8) {
	e.int64(int64(i))
}

func (e *structEncoder) string(s string) {
	e.quote(s)
}

func (e *structEncoder) uint64(ui uint64) {
	encoded := strconv.AppendUint(e.scratch[:0], ui, 10)
	e.Write(encoded)
}

func (e *structEncoder) uint32(ui uint32) {
	e.uint64(uint64(ui))
}

func (e *structEncoder) uint16(ui uint16) {
	e.uint64(uint64(ui))
}

func (e *structEncoder) uint8(ui uint8) {
	e.uint64(uint64(ui))
}

func (e *structEncoder) uint(ui uint) {
	e.uint64(uint64(ui))
}

func (e *structEncoder) quote(s string) {
	e.WriteByte('"')
	e.WriteString(s)
	e.WriteByte('"')
}

func (e *structEncoder) comma() {
	e.WriteByte(',')
}

func (e *structEncoder) attr(name string) {
	e.quote(name)
	e.WriteByte(':')
}

func (e *structEncoder) openBrace() {
	e.WriteByte('{')
}

func (e *structEncoder) closeBrace() {
	e.WriteByte('}')
}
