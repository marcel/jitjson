package fixtures

import (
	"time"
)

type Album struct {
	Title        string  `json:"title"`
	Artist       Artist  `json:"artist"`
	Tracks       []Track `json:"tracks"`
	IsDownloaded bool    `json:"is_downloaded"`
}

type Artist struct {
	Name        string   `json:"name"`
	AlsoKnownAs []string `json:"aka"`
}

type Track struct {
	Number   int           `json:"number"`
	Title    string        `json:"title"`
	Duration time.Duration `json:"duration"`
}

type StructWithNoJSONTags struct {
	FieldWithNoTags string
}
