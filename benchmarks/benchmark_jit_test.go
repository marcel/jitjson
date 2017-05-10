package jitjson

import (
	"encoding/json"
	"testing"

	"github.com/marcel/jitjson/fixtures/media"
	"github.com/marcel/jitjson/fixtures/media/iso"
	"github.com/marcel/jitjson/fixtures/navigation"
	"github.com/stretchr/testify/assert"
)

func TestJSONJitMedia(t *testing.T) {
	album := media.ExampleAlbum
	album.PopularityByMarket = make(map[iso.ISO31661]media.Popularity)
	album.PopularityByMarket[iso.US] = 99 // Change map to just one pair since order is undefined
	r, err := json.Marshal(&album)
	assert.Nil(t, err)

	albumJSON, err := album.MarshalJSON()
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
