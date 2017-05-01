package jitjson

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/marcel/jitjson/fixtures"
	"github.com/stretchr/testify/assert"
)

// func TestJSONStructSearch(t *testing.T) {
// 	search := JSONStructSearch{}

// 	file := "benchmark_test.go"

// 	data, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// structs := search.Structs(data)

// 	// for _, typeSpec := range structs {

// 	// }
// }

// type AlbumEncoder struct {
// 	buf bytes.Buffer
// }

// func (e *AlbumEncoder) encode(album *Album) []byte {
// 	e.buf.Reset()
// 	e.buf.WriteByte('{')
// 	e.buf.WriteString(`"title":"`)
// 	e.buf.WriteString(album.Title)
// 	e.buf.WriteString(`",`)
// 	e.buf.WriteString(`"artist":`)
// 	e.buf.Write(new(ArtistEncoder).encode(&album.Artist))
// 	e.buf.WriteByte(',')
// 	e.buf.WriteString(`"tracks":[`)
// 	trackEncoder := new(TrackEncoder)

// 	lastIndex := len(album.Tracks) - 1
// 	for i, track := range album.Tracks {
// 		e.buf.Write(trackEncoder.encode(&track))
// 		if i != lastIndex {
// 			e.buf.WriteByte(',')
// 		}
// 	}

// 	e.buf.WriteByte(']')
// 	e.buf.WriteByte('}')

// 	return e.buf.Bytes()
// }

// type ArtistEncoder struct {
// 	buf bytes.Buffer
// }

// func (e *ArtistEncoder) encode(artist *Artist) []byte {
// 	e.buf.Reset()

// 	e.buf.WriteByte('{')
// 	e.buf.WriteString(`"name":"`)
// 	e.buf.WriteString(artist.Name)
// 	e.buf.WriteByte('"')
// 	e.buf.WriteByte('}')
// 	return e.buf.Bytes()
// }

// type TrackEncoder struct {
// 	buf bytes.Buffer
// }

// func (e *TrackEncoder) encode(track *Track) []byte {
// 	e.buf.Reset()
// 	e.buf.WriteByte('{')

// 	e.buf.WriteString(`"number":`)
// 	e.buf.WriteString(strconv.FormatInt(int64(track.Number), 10))

// 	e.buf.WriteByte(',')

// 	e.buf.WriteString(`"title":"`)
// 	e.buf.WriteString(track.Title)
// 	e.buf.WriteByte('"')

// 	e.buf.WriteByte(',')

// 	e.buf.WriteString(`"duration":`)
// 	e.buf.WriteString(strconv.Itoa(int(track.Duration.Nanoseconds())))

// 	e.buf.WriteByte('}')

// 	return e.buf.Bytes()
// }

var jsonWithReflectionResult string

func parseDurationStrict(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}

	return duration
}

var album = Album{
	Title: "Some Album Title",
	Artist: Artist{
		Name:        "Some Artist",
		AlsoKnownAs: []string{"Artist Formerly Known as Some Artist"},
	},
	Tracks: []Track{
		Track{Number: 1, Title: "Track 1", Duration: parseDurationStrict("2m30s")},
		Track{Number: 2, Title: "Track 2", Duration: parseDurationStrict("4m22s")},
	},
	IsDownloaded: true,
}

func BenchmarkJSONWithReflection(b *testing.B) {
	var result []byte

	for i := 0; i < b.N; i++ {
		r, err := json.Marshal(&album)
		if err != nil {
			b.Error(err)
		}

		result = r
	}

	jsonWithReflectionResult = string(result)
}

var jsonJittedResult string
var encoder = new(structEncoder)

func TestJSONJit(t *testing.T) {
	r, err := json.Marshal(&album)
	assert.Nil(t, err)

	encoder.albumStruct(album)
	assert.Equal(t, string(r), string(encoder.Bytes()))
}

func BenchmarkJSONJit(b *testing.B) {
	var result []byte

	for i := 0; i < b.N; i++ {
		encoder.Reset()

		encoder.albumStruct(album)
		result = encoder.Bytes()
	}

	jsonJittedResult = string(result)
}
