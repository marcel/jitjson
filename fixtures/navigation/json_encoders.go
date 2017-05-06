package navigation

import "github.com/marcel/jitjson/encoding"

var bufferPool = encoding.NewSyncPool(4096)

type encodingBuffer struct {
	*encoding.Buffer
}

func (s Route) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.Get()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.Put(underlying)
	}()

	buf.routeStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) routeStruct(route Route) {
  e.OpenBrace()

  // "summary":
  e.Write([]byte{0x22, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x3a})
  e.Quote(route.Summary)
  e.Comma()

  // "legs":
  e.Write([]byte{0x22, 0x6c, 0x65, 0x67, 0x73, 0x22, 0x3a})
  e.WriteByte('[')
  for index, element := range route.Legs {
    if index != 0 { e.Comma() }
    e.legStruct(element)
  }
  e.WriteByte(']')

  e.CloseBrace()
}

func (s Leg) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.Get()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.Put(underlying)
	}()

	buf.legStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) legStruct(leg Leg) {
  e.OpenBrace()

  // "distance":
  e.Write([]byte{0x22, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x3a})
  e.Int(leg.Distance)
  e.Comma()

  // "duration":
  e.Write([]byte{0x22, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a})
  e.Int64(int64(leg.Duration))
  e.Comma()

  // "start_address":
  e.Write([]byte{0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3a})
  e.addressStruct(leg.StartAddress)
  e.Comma()

  // "end_address":
  e.Write([]byte{0x22, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3a})
  e.addressStruct(leg.EndAddress)
  e.Comma()

  // "start_location":
  e.Write([]byte{0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a})
  e.locationStruct(leg.StartLocation)
  e.Comma()

  // "end_location":
  e.Write([]byte{0x22, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a})
  e.locationStruct(leg.EndLocation)
  e.Comma()

  // "steps":
  e.Write([]byte{0x22, 0x73, 0x74, 0x65, 0x70, 0x73, 0x22, 0x3a})
  e.WriteByte('[')
  for index, element := range leg.Steps {
    if index != 0 { e.Comma() }
    e.stepStruct(element)
  }
  e.WriteByte(']')

  e.CloseBrace()
}

func (s Step) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.Get()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.Put(underlying)
	}()

	buf.stepStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) stepStruct(step Step) {
  e.OpenBrace()

  // "distance":
  e.Write([]byte{0x22, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x3a})
  e.Int(step.Distance)
  e.Comma()

  // "duration":
  e.Write([]byte{0x22, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a})
  e.Int64(int64(step.Duration))
  e.Comma()

  // "start_location":
  e.Write([]byte{0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a})
  e.locationStruct(step.StartLocation)
  e.Comma()

  // "end_location":
  e.Write([]byte{0x22, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a})
  e.locationStruct(step.EndLocation)
  e.Comma()

  // "travel_mode":
  e.Write([]byte{0x22, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x3a})
  e.Quote(string(step.TravelMode))
  e.Comma()

  // "maneuver":
  e.Write([]byte{0x22, 0x6d, 0x61, 0x6e, 0x65, 0x75, 0x76, 0x65, 0x72, 0x22, 0x3a})
  e.Quote(string(step.Maneuver))
  e.Comma()

  // "instructions":
  e.Write([]byte{0x22, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3a})
  e.Quote(step.Instructions)

  e.CloseBrace()
}

func (s Location) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.Get()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.Put(underlying)
	}()

	buf.locationStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) locationStruct(location Location) {
  e.OpenBrace()

  // "lat":
  e.Write([]byte{0x22, 0x6c, 0x61, 0x74, 0x22, 0x3a})
  e.Quote(location.Lat)
  e.Comma()

  // "lng":
  e.Write([]byte{0x22, 0x6c, 0x6e, 0x67, 0x22, 0x3a})
  e.Quote(location.Lng)

  e.CloseBrace()
}

func (s Address) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.Get()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.Put(underlying)
	}()

	buf.addressStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) addressStruct(address Address) {
  e.OpenBrace()

  // "number":
  e.Write([]byte{0x22, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3a})
  e.Quote(address.Number)
  e.Comma()

  // "street":
  e.Write([]byte{0x22, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x22, 0x3a})
  e.Quote(address.Street)
  e.Comma()

  // "city":
  e.Write([]byte{0x22, 0x63, 0x69, 0x74, 0x79, 0x22, 0x3a})
  e.Quote(address.City)
  e.Comma()

  // "state":
  e.Write([]byte{0x22, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3a})
  e.Quote(address.State)
  e.Comma()

  // "zip_code":
  e.Write([]byte{0x22, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a})
  e.Int(address.ZipCode)
  e.Comma()

  // "country":
  e.Write([]byte{0x22, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x3a})
  e.Quote(address.Country)

  e.CloseBrace()
}

