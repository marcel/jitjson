package media

import "github.com/marcel/jitjson/encoding"

var bufferPool = encoding.NewSyncPool(4096)

type encodingBuffer struct {
	*encoding.Buffer
}

func (s Album) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.Get()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.Put(underlying)
	}()

	buf.albumStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) albumStruct(album Album) {
  e.OpenBrace()

  // "albumType":
  e.Write([]byte{0x22, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3a})
  e.Quote(string(album.AlbumType))
  e.Comma()

  // "artists":
  e.Write([]byte{0x22, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x22, 0x3a})
  e.WriteByte('[')
  for index, element := range album.Artists {
    if index != 0 { e.Comma() }
    e.artistStruct(element)
  }
  e.WriteByte(']')
  e.Comma()

  // "available_markets":
  e.Write([]byte{0x22, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x22, 0x3a})
  e.WriteByte('[')
  for index, element := range album.AvailableMarkets {
    if index != 0 { e.Comma() }
    e.Quote(string(element))
  }
  e.WriteByte(']')
  e.Comma()

  // "genres":
  e.Write([]byte{0x22, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x22, 0x3a})
  e.WriteByte('[')
  for index, element := range album.Genres {
    if index != 0 { e.Comma() }
    e.Quote(string(element))
  }
  e.WriteByte(']')
  e.Comma()

  // "images":
  e.Write([]byte{0x22, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x3a})
  e.WriteByte('[')
  for index, element := range album.Images {
    if index != 0 { e.Comma() }
    e.imageStruct(element)
  }
  e.WriteByte(']')
  e.Comma()

  // "label":
  e.Write([]byte{0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x3a})
  e.Quote(album.Label)
  e.Comma()

  // "name":
  e.Write([]byte{0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a})
  e.Quote(album.Name)
  e.Comma()

  // "popularity":
  e.Write([]byte{0x22, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x22, 0x3a})
  e.Int(int(album.Popularity))
  e.Comma()

  // "release_date":
  e.Write([]byte{0x22, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22, 0x3a})
  jsonBytes, err := album.ReleaseDate.MarshalJSON()
  if err != nil {
    panic(err)
  }
  e.Write(jsonBytes)
  e.Comma()

  // "tracks":
  e.Write([]byte{0x22, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x3a})
  e.WriteByte('[')
  for index, element := range album.Tracks {
    if index != 0 { e.Comma() }
    e.trackStruct(element)
  }
  e.WriteByte(']')

  e.CloseBrace()
}

func (s Image) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.Get()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.Put(underlying)
	}()

	buf.imageStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) imageStruct(image Image) {
  e.OpenBrace()

  // "height":
  e.Write([]byte{0x22, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3a})
  e.Int(image.Height)
  e.Comma()

  // "width":
  e.Write([]byte{0x22, 0x77, 0x69, 0x64, 0x74, 0x68, 0x22, 0x3a})
  e.Int(image.Width)
  e.Comma()

  // "url":
  e.Write([]byte{0x22, 0x75, 0x72, 0x6c, 0x22, 0x3a})
  e.Quote(image.URL)

  e.CloseBrace()
}

func (s Artist) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.Get()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.Put(underlying)
	}()

	buf.artistStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) artistStruct(artist Artist) {
  e.OpenBrace()

  // "genres":
  e.Write([]byte{0x22, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x22, 0x3a})
  e.WriteByte('[')
  for index, element := range artist.Genres {
    if index != 0 { e.Comma() }
    e.Quote(string(element))
  }
  e.WriteByte(']')
  e.Comma()

  // "images":
  e.Write([]byte{0x22, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x3a})
  e.WriteByte('[')
  for index, element := range artist.Images {
    if index != 0 { e.Comma() }
    e.imageStruct(element)
  }
  e.WriteByte(']')
  e.Comma()

  // "popularity":
  e.Write([]byte{0x22, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x22, 0x3a})
  e.Quote(artist.Name)
  e.Comma()

  // "popularity":
  e.Write([]byte{0x22, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x22, 0x3a})
  e.Int(int(artist.Popularity))

  e.CloseBrace()
}

func (s Track) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.Get()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.Put(underlying)
	}()

	buf.trackStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) trackStruct(track Track) {
  e.OpenBrace()

  // "duration":
  e.Write([]byte{0x22, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a})
  e.Int64(int64(track.Duration))
  e.Comma()

  // "is_explicit":
  e.Write([]byte{0x22, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x22, 0x3a})
  e.Bool(bool(track.IsExplicit))
  e.Comma()

  // "is_playable":
  e.Write([]byte{0x22, 0x69, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x3a})
  e.Bool(bool(track.IsPlayable))
  e.Comma()

  // "number":
  e.Write([]byte{0x22, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3a})
  e.Int(track.Number)
  e.Comma()

  // "preview_url":
  e.Write([]byte{0x22, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x3a})
  e.Quote(track.PreviewURL)
  e.Comma()

  // "popularity":
  e.Write([]byte{0x22, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x22, 0x3a})
  e.Int(int(track.Popularity))

  e.CloseBrace()
}

