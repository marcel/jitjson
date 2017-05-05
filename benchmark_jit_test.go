package jitjson

import (
	"encoding/json"
	"testing"

	"github.com/marcel/jitjson/fixtures/navigation"
	"github.com/stretchr/testify/assert"
)

func TestJSONJit(t *testing.T) {
	r, err := json.Marshal(&navigation.ExampleRoute)
	assert.Nil(t, err)

	routeJSON, err := navigation.ExampleRoute.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, string(r), string(routeJSON))
}

func BenchmarkJSONWithReflection(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {
		result = generateReflectionJSON(b)
	}

	_ = result
}

func generateReflectionJSON(b *testing.B) []byte {
	bytes, err := json.Marshal(&navigation.ExampleRoute)
	if err != nil {
		b.Error(err)
	}

	return bytes
}

func BenchmarkJSONJit(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {
		result = generateJitJSON(b)
	}

	_ = result
}

func generateJitJSON(b *testing.B) []byte {
	bytes, err := navigation.ExampleRoute.MarshalJSON()
	if err != nil {
		b.Error(err)
	}

	return bytes
}

// func BenchmarkJSONJitParallel(b *testing.B) {
// 	var result []byte

// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			result = generateJitJSON(b)
// 		}
// 	})

// 	_ = result
// }
