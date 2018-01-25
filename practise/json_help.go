package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func test1() {
	var movies = []Movie{
		{
			Title: "Casablanca",
			Year:  1942,
			Color: false,
			Actors: []string{
				"Humphrey Bogart", "Ingrid Bergman",
			},
		},
		{
			Title: "Cool Hand Luke",
			Year:  1967,
			Color: true,
			Actors: []string{
				"Paul Newman",
			},
		},
		{
			Title: "Bullitt",
			Year:  1968,
			Color: true,
			Actors: []string{
				"Steve McQueen", "Jacqueline Bisset",
			},
		},
		// ...
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	data1, err1 := json.MarshalIndent(movies, "", "    ")
	if err1 != nil {
		log.Fatalf("JSON marshaling failed: %s", err1)
	}
	fmt.Printf("%s\n", data1)
}

func main() {
	test1()
}
