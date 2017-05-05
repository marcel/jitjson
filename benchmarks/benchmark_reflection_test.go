package jitjson

import (
	"encoding/json"
	"testing"

	"github.com/marcel/jitjson/fixtures/media"
	"github.com/marcel/jitjson/fixtures/navigation"
)

func BenchmarkJSONReflectionMedia(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {
		bytes, err := json.Marshal(&media.ExampleAlbum)
		if err != nil {
			b.Error(err)
		}

		result = bytes
	}

	_ = result
}

func BenchmarkJSONReflectionNav(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {
		bytes, err := json.Marshal(&navigation.ExampleRoute)
		if err != nil {
			b.Error(err)
		}

		result = bytes
	}

	_ = result
}
