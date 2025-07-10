package json

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int `json:"released"`
	Actors []string
}

func Marshal() {
	movie := Movie{Title: "Titanic", Year: 2000, Actors: []string{"Tom Hanks", "Tilda Swinton"}}
	data, err := json.MarshalIndent(movie, "", "  ")
	if err != nil {
		fmt.Printf("json error%#v", err)
		return
	}
	fmt.Printf("%s\n", data)
}

func Unmarshal() {
	j := `{"Title": "test", "released": 1958, "Actors": ["jsmae", "jasdfkja"]}`

	var titles Movie
	if err := json.Unmarshal([]byte(j), &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
