package media

import (
	"time"
)

type AlbumType string

const (
	AlbumAlbumType       AlbumType = "album"
	SingleAlbumType      AlbumType = "single"
	CompilationAlbumType AlbumType = "compilation"
)

type ISO31661Alpha2 string

var AllISO31661Alpha2 = struct {
	AD, AE /* etc */, UK, US ISO31661Alpha2
}{
	AD: "AD", AE: "AE" /* etc */, UK: "UK", US: "US",
}

type Genre string

const (
	ProgRock   Genre = "Prog Rock"
	PostGrunge Genre = "Post-Grunge"
)

type Popularity int

type Album struct {
	AlbumType        AlbumType        `json:"album_type"`
	Artists          []Artist         `json:"artists"`
	AvailableMarkets []ISO31661Alpha2 `json:"available_markets"`
	Genres           []Genre          `json:"genres"`
	Images           []Image          `json:"images"`
	Label            string           `json:"label"`
	Name             string           `json:"name"`
	Popularity       Popularity       `json:"popularity"`
	ReleaseDate      time.Time        `json:"release_date"`
	Tracks           []Track          `json:"tracks"`
}

var ExampleAlbum = Album{
	AlbumType: AlbumAlbumType,
	Artists:   []Artist{ExampleArtist},
	AvailableMarkets: []ISO31661Alpha2{
		AllISO31661Alpha2.US,
		AllISO31661Alpha2.UK,
	},
	Genres:      []Genre{PostGrunge, ProgRock},
	Images:      []Image{ExampleImage},
	Label:       "Unstable",
	Name:        "Some Album Title",
	Popularity:  14,
	ReleaseDate: time.Now().AddDate(-5, 0, 0),
	Tracks:      []Track{ExampleTrack},
}

type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

var ExampleImage = Image{
	Height: 1000,
	Width:  1000,
	URL:    "http://www.billboard.com/files/styles/900_wide/public/media/Joy-Division-Unknown-Pleasures-album-covers-billboard-1000x1000.jpg",
}

type Artist struct {
	Genres     []Genre    `json:"genres"`
	Images     []Image    `json:"images"`
	Name       string     `json:"popularity"`
	Popularity Popularity `json:"popularity"`
}

var ExampleArtist = Artist{
	Genres:     []Genre{ProgRock, PostGrunge},
	Images:     []Image{ExampleImage},
	Name:       "Some Artist's Name",
	Popularity: 42,
}

type Track struct {
	Duration   time.Duration `json:"duration"`
	IsExplicit bool          `json:"is_explicit"`
	IsPlayable bool          `json:"is_playable"`
	Number     int           `json:"number"`
	PreviewURL string        `json:"preview_url"`
	Popularity Popularity    `json:"popularity"`
}

var ExampleTrack = Track{
	Duration:   time.Duration(256 * time.Second),
	IsExplicit: false,
	IsPlayable: true,
	Number:     1,
	PreviewURL: "http://Example.com/preview.ogg",
	Popularity: 81,
}

type StructWithNoJSONTags struct {
	FieldWithNoTags string
}
