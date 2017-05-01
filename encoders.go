package jitjson

import . "github.com/marcel/jitjson/fixtures"

func (e *structEncoder) albumStruct(album Album) {
	e.openBrace()

	e.attr("title")
	e.string(album.Title)
	e.comma()

	e.attr("artist")
	e.artistStruct(album.Artist)
	e.comma()

	e.attr("tracks")
	e.WriteByte('[')
	for index, element := range album.Tracks {
		if index != 0 {
			e.comma()
		}
		e.trackStruct(element)
	}
	e.WriteByte(']')
	e.comma()

	e.attr("is_downloaded")
	e.bool(album.IsDownloaded)

	e.closeBrace()
}

func (e *structEncoder) artistStruct(artist Artist) {
	e.openBrace()

	e.attr("name")
	e.string(artist.Name)
	e.comma()

	e.attr("aka")
	e.WriteByte('[')
	for index, element := range artist.AlsoKnownAs {
		if index != 0 {
			e.comma()
		}
		e.string(element)
	}
	e.WriteByte(']')

	e.closeBrace()
}

func (e *structEncoder) trackStruct(track Track) {
	e.openBrace()

	e.attr("number")
	e.int(int(track.Number))
	e.comma()

	e.attr("title")
	e.string(track.Title)
	e.comma()

	e.attr("duration")
	e.int64(int64(track.Duration))

	e.closeBrace()
}
