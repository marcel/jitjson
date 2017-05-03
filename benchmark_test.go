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

	routeJSON, err := (&navigation.ExampleRoute).MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, string(r), string(routeJSON))
}

func BenchmarkJSONWithReflection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&navigation.ExampleRoute)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJSONJit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := (&navigation.ExampleRoute).MarshalJSON()
		if err != nil {
			b.Error(err)
		}
	}
}
