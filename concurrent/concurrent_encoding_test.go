package concurrent

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	cb := new(ConcurrentBuffer)
	cb.Buffer = new(bytes.Buffer)
	cb.Encode(&ExampleGarage)
	fmt.Println(cb.String())

	fmt.Println("")
	bytes, _ := ExampleGarage.MarshalJSON()
	fmt.Println(string(bytes))
}

func BenchmarkSerial(b *testing.B) {
	var results []byte

	for i := 0; i < b.N; i++ {
		bytes, _ := ExampleGarage.MarshalJSON()
		results = bytes
	}

	_ = results
}

func BenchmarkConcurrent(b *testing.B) {
	var results []byte

	for i := 0; i < b.N; i++ {
		cb := new(ConcurrentBuffer)
		cb.Buffer = new(bytes.Buffer)
		cb.Encode(&ExampleGarage)
		results = cb.Bytes()
	}

	_ = results
}
