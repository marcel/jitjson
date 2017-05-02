package jitjson

import (
	"encoding/json"
	"testing"

	"github.com/marcel/jitjson/fixtures/media"
	"github.com/stretchr/testify/assert"
)

var jsonWithReflectionResult string

func BenchmarkJSONWithReflection(b *testing.B) {
	var result []byte

	for i := 0; i < b.N; i++ {
		r, err := json.Marshal(&media.ExampleAlbum)
		if err != nil {
			b.Error(err)
		}

		result = r
	}

	jsonWithReflectionResult = string(result)
}

var jsonJittedResult string

func TestJSONJit(t *testing.T) {
	r, err := json.Marshal(&media.ExampleAlbum)
	assert.Nil(t, err)

	albumJSON, err := (&media.ExampleAlbum).MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, string(r), string(albumJSON))
}

func BenchmarkJSONJit(b *testing.B) {
	var result []byte

	for i := 0; i < b.N; i++ {
		albumJSON, _ := (&media.ExampleAlbum).MarshalJSON()
		result = albumJSON
	}

	jsonJittedResult = string(result)
}
