package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type InsultStruct struct {
	Number string `json:"number"` // The api sends the insult number as a string...
	Insult string `json:"insult"`
}

func main() {
	resp, err := http.Get("https://evilinsult.com/generate_insult.php?lang=en&type=json")
	if err != nil {
		log.Fatalln(err)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert json to go struct
	var insult InsultStruct
	err = json.Unmarshal(body, &insult)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(`"%s" [Insult #%s]`, insult.Insult, insult.Number)
}
