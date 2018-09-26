package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

const defaultFileName = "w.json"

func main() {
	var fileName = defaultFileName
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	fmt.Println(fileName)

	var w Wheel
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v\n", w)

	var data, err = json.MarshalIndent(w, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var result Wheel

	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%#v\n", result)
}
