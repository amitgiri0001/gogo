package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	jsonVal := `
	[
		{
			"first_name": "Amit",
			"last_name": "Giri"
		},
		{
			"first_name": "Aman",
			"last_name": "Giri"
		}
	]
   `

	structVar := []Person{}
	err := json.Unmarshal([]byte(jsonVal), &structVar)

	if err != nil {
		log.Println("Error in unmarshalling json", err)
	}

	log.Printf("Unmarshelled: %+v", structVar)

	jsonVal2, err := json.Marshal(structVar)
	log.Printf("Marshelled: %+v", string(jsonVal2))

}
