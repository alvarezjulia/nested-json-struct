package main

import (
	"encoding/json"
	"log"
)

const jsonText = `
{
  "libraries": [
    {
      "name": "Katakrak",
      "city": "Pamplona",
      "country": "Spain",
      "books": [
        {
          "title": "Harry Potter and the Sorcerer's Stone",
          "author": "J. K. Rowling",
          "availability": true
        }
      ]
    },
    {
      "name": "Re-Read Librería Lowcost",
      "city": "Pamplona",
      "country": "Spain",
      "books": [
        {
          "title": "The Subtle Art of Not Giving a F*ck",
          "author": "Mark Manson",
          "availability": false
        }
      ]
    }
  ]
}`

type Libraries struct {
	LibrariesList []struct {
		Name    string `json:"name"`
		City    string `json:"city"`
		Country string `json:"country"`
		Books   []struct {
			Title        string `json:"title"`
			Author       string `json:"author"`
			Availability bool   `json:"availability"`
		} `json:"books"`
	} `json:"libraries"`
}

func main() {
	var librariesInformation Libraries
	err := json.Unmarshal([]byte(jsonText), &librariesInformation)
	if err != nil {
		log.Fatal("error unmarshaling json: ", err)
	}
	log.Printf("librariesInformation: %+v", librariesInformation)
}
