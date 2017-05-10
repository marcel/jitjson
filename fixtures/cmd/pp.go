package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/marcel/jitjson/fixtures/media"
	"github.com/marcel/jitjson/fixtures/navigation"
)

func main() {
	example := flag.String("e", "nav", "which json to print [nav | media]")
	method := flag.String("m", "jit", "whether to use jit or reflection [jit | r]")

	flag.Parse()

	var jsonBytes []byte
	var err error

	switch *example {
	case "nav":
		switch *method {
		case "jit":
			jsonBytes, err = navigation.ExampleRoute.MarshalJSON()
		case "r":
			jsonBytes, err = json.Marshal(navigation.ExampleRoute)
		}
	case "media":
		switch *method {
		case "jit":
			jsonBytes, err = media.ExampleAlbum.MarshalJSON()
		case "r":
			jsonBytes, err = json.Marshal(media.ExampleAlbum)
		}
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))
}
