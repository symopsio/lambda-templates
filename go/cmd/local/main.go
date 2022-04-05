package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	i "github.com/symopsio/lambda-templates/go/internal"
)

// main lets you test a Sym event payload locally
func main() {
	var e, d bool
	flag.BoolVar(&e, "e", false, "escalate")
	flag.BoolVar(&d, "d", false, "deescalate")
	flag.Parse()
	if (e && d) || !(e || d) {
		log.Fatal("must supply either -e or -d")
	}
	file := "escalate"
	if d {
		file = "deescalate"
	}
	event, err := loadEvent(file)
	if err != nil {
		log.Fatal(err)
	}

	result, err := i.HandleSymEvent(*event)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Got result: %+v\n", result)
}

// loadEvent reads a sym event from a local test file
func loadEvent(fileName string) (*i.SymEvent, error) {
	path := fmt.Sprintf("../../../test/%s.json", fileName)
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var event i.SymEvent
	json.Unmarshal(byteValue, &event)
	return &event, nil
}
