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
	AlbumType        AlbumType        `json:"albumType"`
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
	Artists: []Artist{
		Artist{
			Genres: []Genre{ProgRock, PostGrunge},
			Images: []Image{
				Image{
					Height: 1000,
					Width:  1000,
					URL:    "http://www.billboard.com/files/styles/900_wide/public/media/Joy-Division-Unknown-Pleasures-album-covers-billboard-1000x1000.jpg",
				},
			},
			Name:       "Some Artist's Name",
			Popularity: 42,
		},
	},
	AvailableMarkets: []ISO31661Alpha2{
		AllISO31661Alpha2.US,
		AllISO31661Alpha2.UK,
	},
	Genres: []Genre{PostGrunge, ProgRock},
	Images: []Image{
		Image{
			Height: 1000,
			Width:  1000,
			URL:    "http://www.billboard.com/files/styles/900_wide/public/media/Joy-Division-Unknown-Pleasures-album-covers-billboard-1000x1000.jpg",
		},
	},
	Label:       "Unstable",
	Name:        "Some Album Title",
	Popularity:  14,
	ReleaseDate: time.Now().AddDate(-5, 0, 0),
	Tracks: []Track{
		Track{
			Title:      "Track 1",
			Duration:   time.Duration(256 * time.Second),
			IsExplicit: false,
			IsPlayable: true,
			Number:     1,
			PreviewURL: "http://audiogalaxy.com/preview.ogg",
			Popularity: 81,
		},
		Track{
			Title:      "Track 2",
			Duration:   time.Duration(342 * time.Second),
			IsExplicit: true,
			IsPlayable: true,
			Number:     2,
			PreviewURL: "http://audiogalaxy.com/preview.ogg",
			Popularity: 99,
		},
		Track{
			Title:      "Track 3",
			Duration:   time.Duration(213 * time.Second),
			IsExplicit: false,
			IsPlayable: true,
			Number:     3,
			PreviewURL: "http://audiogalaxy.com/preview.ogg",
			Popularity: 87,
		},
		Track{
			Title:      "Track 4",
			Duration:   time.Duration(145 * time.Second),
			IsExplicit: false,
			IsPlayable: true,
			Number:     4,
			PreviewURL: "http://audiogalaxy.com/preview.ogg",
			Popularity: 75,
		},
		Track{
			Title:      "Track 5",
			Duration:   time.Duration(321 * time.Second),
			IsExplicit: false,
			IsPlayable: true,
			Number:     5,
			PreviewURL: "http://audiogalaxy.com/preview.ogg",
			Popularity: 81,
		},
		Track{
			Title:      "Track 6",
			Duration:   time.Duration(387 * time.Second),
			IsExplicit: false,
			IsPlayable: true,
			Number:     6,
			PreviewURL: "http://audiogalaxy.com/preview.ogg",
			Popularity: 88,
		},
		Track{
			Title:      "Track 7",
			Duration:   time.Duration(327 * time.Second),
			IsExplicit: false,
			IsPlayable: true,
			Number:     7,
			PreviewURL: "http://audiogalaxy.com/preview.ogg",
			Popularity: 82,
		},
		Track{
			Title:      "Track 8",
			Duration:   time.Duration(287 * time.Second),
			IsExplicit: false,
			IsPlayable: true,
			Number:     8,
			PreviewURL: "http://audiogalaxy.com/preview.ogg",
			Popularity: 78,
		},
		Track{
			Title:      "Track 9",
			Duration:   time.Duration(382 * time.Second),
			IsExplicit: false,
			IsPlayable: true,
			Number:     9,
			PreviewURL: "http://audiogalaxy.com/preview.ogg",
			Popularity: 98,
		},
		Track{
			Title:      "Track 10",
			Duration:   time.Duration(387 * time.Second),
			IsExplicit: false,
			IsPlayable: true,
			Number:     10,
			PreviewURL: "http://audiogalaxy.com/preview.ogg",
			Popularity: 68,
		},
	},
}

type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

type Artist struct {
	Genres     []Genre `json:"genres"`
	Images     []Image `json:"images"`
	Name       string  `json:"popularity"`
	Popularity `json:"popularity"`
}

type Track struct {
	Title      string
	Duration   time.Duration `json:"duration"`
	IsExplicit bool          `json:"is_explicit"`
	IsPlayable bool          `json:"is_playable"`
	Number     int           `json:"number"`
	PreviewURL string        `json:"preview_url"`
	Popularity Popularity    `json:"popularity"`
}

type StructWithNoJSONTags struct {
	FieldWithNoTags string
}
