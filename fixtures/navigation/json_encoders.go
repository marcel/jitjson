package navigation

import "github.com/marcel/jitjson/encoding"

type encodingBuffer struct {
	encoding.Buffer
}

func (s Route) MarshalJSON() ([]byte, error) {
	buf := encodingBuffer{}
	buf.routeStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) routeStruct(route Route) {
  e.OpenBrace()

  e.Attr("summary")
  e.String(string(route.Summary))
  e.Comma()

  e.Attr("legs")
  e.WriteByte('[')
  for index, element := range route.Legs {
    if index != 0 { e.Comma() }
    e.legStruct(element)
  }
  e.WriteByte(']')

  e.CloseBrace()
}

func (s Leg) MarshalJSON() ([]byte, error) {
	buf := encodingBuffer{}
	buf.legStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) legStruct(leg Leg) {
  e.OpenBrace()

  e.Attr("distance")
  e.Int(int(leg.Distance))
  e.Comma()

  e.Attr("duration")
  e.Int64(int64(leg.Duration))
  e.Comma()

  e.Attr("start_address")
  e.addressStruct(leg.StartAddress)
  e.Comma()

  e.Attr("end_address")
  e.addressStruct(leg.EndAddress)
  e.Comma()

  e.Attr("start_location")
  e.locationStruct(leg.StartLocation)
  e.Comma()

  e.Attr("end_location")
  e.locationStruct(leg.EndLocation)
  e.Comma()

  e.Attr("steps")
  e.WriteByte('[')
  for index, element := range leg.Steps {
    if index != 0 { e.Comma() }
    e.stepStruct(element)
  }
  e.WriteByte(']')

  e.CloseBrace()
}

func (s Step) MarshalJSON() ([]byte, error) {
	buf := encodingBuffer{}
	buf.stepStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) stepStruct(step Step) {
  e.OpenBrace()

  e.Attr("distance")
  e.Int(int(step.Distance))
  e.Comma()

  e.Attr("duration")
  e.Int64(int64(step.Duration))
  e.Comma()

  e.Attr("start_location")
  e.locationStruct(step.StartLocation)
  e.Comma()

  e.Attr("end_location")
  e.locationStruct(step.EndLocation)
  e.Comma()

  e.Attr("travel_mode")
  e.String(string(step.TravelMode))
  e.Comma()

  e.Attr("maneuver")
  e.String(string(step.Maneuver))
  e.Comma()

  e.Attr("instructions")
  e.String(string(step.Instructions))

  e.CloseBrace()
}

func (s Location) MarshalJSON() ([]byte, error) {
	buf := encodingBuffer{}
	buf.locationStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) locationStruct(location Location) {
  e.OpenBrace()

  e.Attr("lat")
  e.String(string(location.Lat))
  e.Comma()

  e.Attr("lng")
  e.String(string(location.Lng))

  e.CloseBrace()
}

func (s Address) MarshalJSON() ([]byte, error) {
	buf := encodingBuffer{}
	buf.addressStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) addressStruct(address Address) {
  e.OpenBrace()

  e.Attr("number")
  e.String(string(address.Number))
  e.Comma()

  e.Attr("street")
  e.String(string(address.Street))
  e.Comma()

  e.Attr("city")
  e.String(string(address.City))
  e.Comma()

  e.Attr("state")
  e.String(string(address.State))
  e.Comma()

  e.Attr("zip_code")
  e.Int(int(address.ZipCode))
  e.Comma()

  e.Attr("country")
  e.String(string(address.Country))

  e.CloseBrace()
}

