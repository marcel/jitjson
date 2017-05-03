package media

import "github.com/marcel/jitjson/encoding"

var bufferPool = encoding.NewSyncPool(4096)

type encodingBuffer struct {
	*encoding.Buffer
}

func (s Album) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.GetBuffer()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.PutBuffer(underlying)
	}()

	buf.albumStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) albumStruct(album Album) {
  e.OpenBrace()

  e.Attr("albumType")
  e.String(string(album.AlbumType))
  e.Comma()

  e.Attr("artists")
  e.WriteByte('[')
  for index, element := range album.Artists {
    if index != 0 { e.Comma() }
    e.artistStruct(element)
  }
  e.WriteByte(']')
  e.Comma()

  e.Attr("available_markets")
  e.WriteByte('[')
  for index, element := range album.AvailableMarkets {
    if index != 0 { e.Comma() }
    e.String(string(element))
  }
  e.WriteByte(']')
  e.Comma()

  e.Attr("genres")
  e.WriteByte('[')
  for index, element := range album.Genres {
    if index != 0 { e.Comma() }
    e.String(string(element))
  }
  e.WriteByte(']')
  e.Comma()

  e.Attr("images")
  e.WriteByte('[')
  for index, element := range album.Images {
    if index != 0 { e.Comma() }
    e.imageStruct(element)
  }
  e.WriteByte(']')
  e.Comma()

  e.Attr("label")
  e.String(string(album.Label))
  e.Comma()

  e.Attr("name")
  e.String(string(album.Name))
  e.Comma()

  e.Attr("popularity")
  e.Int(int(album.Popularity))
  e.Comma()

  e.Attr("release_date")
  jsonBytes, err := album.ReleaseDate.MarshalJSON()
  if err != nil {
    panic(err)
  }
  e.Write(jsonBytes)
  e.Comma()

  e.Attr("tracks")
  e.WriteByte('[')
  for index, element := range album.Tracks {
    if index != 0 { e.Comma() }
    e.trackStruct(element)
  }
  e.WriteByte(']')

  e.CloseBrace()
}

func (s Image) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.GetBuffer()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.PutBuffer(underlying)
	}()

	buf.imageStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) imageStruct(image Image) {
  e.OpenBrace()

  e.Attr("height")
  e.Int(int(image.Height))
  e.Comma()

  e.Attr("width")
  e.Int(int(image.Width))
  e.Comma()

  e.Attr("url")
  e.String(string(image.URL))

  e.CloseBrace()
}

func (s Artist) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.GetBuffer()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.PutBuffer(underlying)
	}()

	buf.artistStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) artistStruct(artist Artist) {
  e.OpenBrace()

  e.Attr("genres")
  e.WriteByte('[')
  for index, element := range artist.Genres {
    if index != 0 { e.Comma() }
    e.String(string(element))
  }
  e.WriteByte(']')
  e.Comma()

  e.Attr("images")
  e.WriteByte('[')
  for index, element := range artist.Images {
    if index != 0 { e.Comma() }
    e.imageStruct(element)
  }
  e.WriteByte(']')
  e.Comma()

  e.Attr("popularity")
  e.String(string(artist.Name))
  e.Comma()

  e.Attr("popularity")
  e.Int(int(artist.Popularity))

  e.CloseBrace()
}

func (s Track) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.GetBuffer()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.PutBuffer(underlying)
	}()

	buf.trackStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) trackStruct(track Track) {
  e.OpenBrace()

  e.Attr("duration")
  e.Int64(int64(track.Duration))
  e.Comma()

  e.Attr("is_explicit")
  e.Bool(bool(track.IsExplicit))
  e.Comma()

  e.Attr("is_playable")
  e.Bool(bool(track.IsPlayable))
  e.Comma()

  e.Attr("number")
  e.Int(int(track.Number))
  e.Comma()

  e.Attr("preview_url")
  e.String(string(track.PreviewURL))
  e.Comma()

  e.Attr("popularity")
  e.Int(int(track.Popularity))

  e.CloseBrace()
}

