package concurrent

import (
	"bytes"
	"sync"

	"github.com/marcel/jitjson/encoding"
)

type Garage struct {
	Cars              []Car `json:"cars"`
	TotalChargers     int   `json:"total_chargers"`
	ChargersAvailable int   `json:"chargers_available"`
	TotalSpots        int   `json:"total_spots"`
	SpotsAvailable    int   `json:"spots_available"`
}

type Make string

const (
	Tesla Make = "Tesla"
	Chevy Make = "Chevy"
)

type Model string

const (
	ModelS Model = "Model S"
	ModelX Model = "Model X"
	Volt   Model = "Volt"
)

type Color string

const (
	Red   Color = "red"
	Black Color = "black"
	Grey  Color = "grey"
	White Color = "white"
)

type Car struct {
	Color `json:"color"`
	Make  `json:"make"`
	Model `json:"model"`
	Year  int `json:"year"`
	VIN   int `json:"vin"`
}

var ExampleGarage = Garage{
	TotalChargers:     6,
	ChargersAvailable: 3,
	TotalSpots:        10,
	SpotsAvailable:    3,
	Cars: []Car{
		Car{Red, Tesla, ModelS, 2016, 98765678567},
		Car{Black, Tesla, ModelX, 2017, 9978678567},
		Car{Grey, Tesla, ModelS, 2014, 678678567},
		Car{Red, Chevy, Volt, 2017, 98765678567},
		Car{White, Tesla, ModelS, 2016, 98765678567},
		Car{Red, Tesla, ModelS, 2016, 98765678567},
		Car{Black, Tesla, ModelX, 2017, 9978678567},
		Car{Grey, Tesla, ModelS, 2014, 678678567},
		Car{Red, Chevy, Volt, 2017, 98765678567},
		Car{White, Tesla, ModelS, 2016, 98765678567},
		Car{Red, Tesla, ModelS, 2016, 98765678567},
		Car{Black, Tesla, ModelX, 2017, 9978678567},
		Car{Grey, Tesla, ModelS, 2014, 678678567},
		Car{Red, Chevy, Volt, 2017, 98765678567},
		Car{White, Tesla, ModelS, 2016, 98765678567},
	},
}

type CoalescingConcurrentBuffer struct {
	Size    int
	Buffers []*bytes.Buffer
}

type ConcurrentBuffer struct {
	*bytes.Buffer
}

func (cb *ConcurrentBuffer) Encode(garage *Garage) {
	primitivesBuffer := bufferPool.GetBuffer()
	defer func() {
		primitivesBuffer.Reset()
		bufferPool.PutBuffer(primitivesBuffer)
	}()

	allCarsBuf := bufferPool.GetBuffer()
	defer func() {
		allCarsBuf.Reset()
		bufferPool.PutBuffer(allCarsBuf)
	}()

	buffers := []*encoding.Buffer{primitivesBuffer, allCarsBuf}

	primitivesWg := sync.WaitGroup{}
	primitivesWg.Add(1)

	go func() {
		defer primitivesWg.Done()
		underlying := bufferPool.GetBuffer()
		buf := encodingBuffer{Buffer: underlying}
		defer func() {
			underlying.Reset()
			bufferPool.PutBuffer(underlying)
		}()

		buf.Attr("total_chargers")
		buf.Int(garage.TotalChargers)
		buf.Comma()

		buf.Attr("chargers_available")
		buf.Int(garage.ChargersAvailable)
		buf.Comma()

		buf.Attr("total_spots")
		buf.Int(garage.TotalSpots)
		buf.Comma()

		buf.Attr("spots_available")
		buf.Int(garage.SpotsAvailable)
		buf.Comma()

		primitivesBuffer.Write(buf.Bytes())
	}()

	carBufferSlots := len(garage.Cars)
	wg := sync.WaitGroup{}
	wg.Add(carBufferSlots)

	carBuffers := make([]*encoding.Buffer, carBufferSlots, carBufferSlots)
	for i := 0; i < carBufferSlots; i++ {
		carBuffers[i] = bufferPool.GetBuffer()
	}

	for i := range garage.Cars {
		go func(index int) {
			defer wg.Done()
			buf := carBuffers[index]
			defer func() {
				buf.Reset()
				bufferPool.PutBuffer(buf)
			}()

			car := garage.Cars[index]
			bytes, _ := car.MarshalJSON()
			buf.Write(bytes)
		}(i)
	}

	wg.Wait()

	allCarsBuf.WriteString(`"cars":`)
	allCarsBuf.WriteByte('[')
	for i, carBuffer := range carBuffers {
		if i != 0 {
			allCarsBuf.WriteByte(',')
		}
		allCarsBuf.Write(carBuffer.Bytes())
	}
	allCarsBuf.WriteByte(']')

	primitivesWg.Wait()
	allBufs := bufferPool.GetBuffer()
	defer func() {
		allBufs.Reset()
		bufferPool.PutBuffer(allBufs)
	}()

	totalSize := 0
	for i := range buffers {
		totalSize += buffers[i].Len()
	}

	allBufs.Grow(totalSize)

	allBufs.WriteByte('{')
	for i := range buffers {
		allBufs.Write(buffers[i].Bytes())
	}
	allBufs.WriteByte('}')

	cb.Write(allBufs.Bytes())
}
