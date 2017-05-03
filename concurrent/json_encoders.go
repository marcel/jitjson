package concurrent

import "github.com/marcel/jitjson/encoding"

var bufferPool = encoding.NewSyncPool(4096)

type encodingBuffer struct {
	*encoding.Buffer
}

func (s Garage) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.GetBuffer()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.PutBuffer(underlying)
	}()

	buf.garageStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) garageStruct(garage Garage) {
  e.OpenBrace()

  e.Attr("cars")
  e.WriteByte('[')
  for index, element := range garage.Cars {
    if index != 0 { e.Comma() }
    e.carStruct(element)
  }
  e.WriteByte(']')
  e.Comma()

  e.Attr("total_chargers")
  e.Int(int(garage.TotalChargers))
  e.Comma()

  e.Attr("chargers_available")
  e.Int(int(garage.ChargersAvailable))
  e.Comma()

  e.Attr("total_spots")
  e.Int(int(garage.TotalSpots))
  e.Comma()

  e.Attr("spots_available")
  e.Int(int(garage.SpotsAvailable))

  e.CloseBrace()
}

func (s Car) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.GetBuffer()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.PutBuffer(underlying)
	}()

	buf.carStruct(s)
	return buf.Bytes(), nil
}

func (e *encodingBuffer) carStruct(car Car) {
  e.OpenBrace()

  e.Attr("color")
  e.String(string(car.Color))
  e.Comma()

  e.Attr("make")
  e.String(string(car.Make))
  e.Comma()

  e.Attr("model")
  e.String(string(car.Model))
  e.Comma()

  e.Attr("year")
  e.Int(int(car.Year))
  e.Comma()

  e.Attr("vin")
  e.Int(int(car.VIN))

  e.CloseBrace()
}

