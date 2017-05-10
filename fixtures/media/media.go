package media

import (
	"time"

	"github.com/marcel/jitjson/fixtures/media/iso"
)

type AlbumType string

const (
	AlbumAlbumType       AlbumType = "album"
	SingleAlbumType      AlbumType = "single"
	CompilationAlbumType AlbumType = "compilation"
)

type Genre string

const (
	ProgRock   Genre = "Prog Rock"
	PostGrunge Genre = "Post-Grunge"
)

type Popularity int

type Album struct {
	AlbumType          AlbumType                   `json:"album_type"`
	Artists            []Artist                    `json:"artists"`
	AvailableMarkets   []iso.ISO31661              `json:"available_markets"`
	Genres             []Genre                     `json:"genres"`
	Images             []Image                     `json:"images"`
	Label              string                      `json:"label"`
	Name               string                      `json:"name"`
	Popularity         Popularity                  `json:"popularity"`
	ReleaseDate        time.Time                   `json:"release_date"`
	Tracks             []Track                     `json:"tracks"`
	PopularityByMarket map[iso.ISO31661]Popularity `json:"popularity_by_market"`
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
					URL:    "http://billboard.com/files/82782376.jpg",
				},
			},
			Name:       "Some Artist's Name",
			Popularity: 42,
		},
	},
	AvailableMarkets: []iso.ISO31661{
		iso.AF,
		iso.AO,
		iso.AT,
		iso.BH,
		iso.BQ,
		iso.BL,
		iso.BT,
		iso.CA,
		iso.CL,
		iso.CX,
		iso.DE,
		iso.DK,
		iso.DO,
		iso.AD,
		iso.AE,
		iso.AX,
		iso.MX,
		iso.NO,
		iso.US,
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
	PopularityByMarket: map[iso.ISO31661]Popularity{
		iso.AF: 98,
		iso.AO: 82,
		iso.AT: 96,
		iso.BH: 86,
		iso.BQ: 92,
		iso.BL: 87,
		iso.BT: 28,
		iso.CA: 98,
		iso.CL: 98,
		iso.CX: 78,
		iso.DE: 76,
		iso.DK: 83,
		iso.DO: 67,
		iso.AD: 33,
		iso.AE: 98,
		iso.AX: 92,
		iso.MX: 97,
		iso.NO: 87,
		iso.US: 99,
	},
}

type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

type Artist struct {
	Genres     []Genre    `json:"genres"`
	Images     []Image    `json:"images"`
	Name       string     `json:"name"`
	Popularity Popularity `json:"popularity"`
}

type Track struct {
	Title      string        `json:"title"`
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
