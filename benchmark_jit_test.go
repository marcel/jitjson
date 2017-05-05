package jitjson

import (
	"encoding/json"
	"testing"

	"github.com/marcel/jitjson/fixtures/media"
	"github.com/marcel/jitjson/fixtures/navigation"
	"github.com/stretchr/testify/assert"
)

func TestJSONJitMedia(t *testing.T) {
	r, err := json.Marshal(&media.ExampleAlbum)
	assert.Nil(t, err)

	albumJSON, err := media.ExampleAlbum.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, string(r), string(albumJSON))
}

func TestJSONJitNav(t *testing.T) {
	r, err := json.Marshal(&navigation.ExampleRoute)
	assert.Nil(t, err)

	routeJSON, err := navigation.ExampleRoute.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, string(r), string(routeJSON))
}

func BenchmarkJSONJitMedia(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {

		bytes, err := media.ExampleAlbum.MarshalJSON()
		if err != nil {
			b.Error(err)
		}
		result = bytes
	}

	_ = result
}

func BenchmarkJSONJitNav(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {

		bytes, err := navigation.ExampleRoute.MarshalJSON()
		if err != nil {
			b.Error(err)
		}
		result = bytes
	}

	_ = result
}

// func BenchmarkJSONJitParallel(b *testing.B) {
// 	var result []byte

// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			result = generateNavJitJSON(b)
// 		}
// 	})

// 	_ = result
// }
