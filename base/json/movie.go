package json

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{"A and B", 1942, false, []string{"A", "B"}},
	{"C and D", 1967, true, []string{"C", "D"}},
	{"E and F", 1972, true, []string{"E", "F"}},
}

func Json1() {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshalling faild:%s\n", err)
	}
	fmt.Printf("%s\n", data)

	var title []struct{
		Title string
	}
	if err := json.Unmarshal(data, &title); err != nil {
		log.Fatalf("JSON unmarshaling failed:%s", err)
	}
	fmt.Println(title)
}
