package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie define an arbitrary movie structure for demo purposes
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergmann"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jaqueline Bisset"}}}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	data, err = json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string } // needs to be uppercase to be imported
	if err = json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("json unmarshaling failed: %s", err)
	} // uses reflection
	fmt.Printf("%s\n", titles)
}
