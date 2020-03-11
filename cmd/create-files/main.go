package main

import (
	"fmt"
	"os"
	"time"

	"sloan.com/service/internal/event"
	"sloan.com/service/internal/platform/utility"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 1 {
		fmt.Printf("\ninvalid args\n")
		os.Exit(1)
	}

	evts := []*event.Event{
		&event.Event{
			Id:   "foo",
			Name: "bar",
			Context: []string{
				"foo",
				"bar",
				"baz",
			},
			Data: map[string]interface{}{
				"foo": map[string]interface{}{
					"baz": "bar",
				},
			},
			Time: time.Now().UnixNano() / 1000, // ms
		},
		&event.Event{
			Id:   "bsdfg",
			Name: "sdfgsgd",
			Context: []string{
				"sdfgsgd",
				"sdfgsdfg",
				"sdfgsgdf",
			},
			Data: map[string]interface{}{
				"fsdfgoo": map[string]interface{}{
					"sdfgsgdf": "sdfgdfg",
				},
			},
			Time: time.Now().UnixNano() / 1000, // ms
		},
	}

	utility.DumpEventsToFile(evts, argsWithoutProg[0])
}
