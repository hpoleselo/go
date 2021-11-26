package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	// Tells Go when we receive/create JSON
	// We can now unmarshal the jsonExample as we have a data structure
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func main() {
	readJSON()
	writeJSON()

}

func writeJSON() {
	var mySlice []Person
	var structToBeMarshelled Person
	structToBeMarshelled.FirstName = "Homer"
	structToBeMarshelled.LastName = "Simpson"
	structToBeMarshelled.Age = 40

	// Since we want the JSON to have any length we store it in a slice as well, so what we'll be converting is the final slice
	mySlice = append(mySlice, structToBeMarshelled)

	var structToBeMarshelled2 Person
	structToBeMarshelled2.FirstName = "Margie"
	structToBeMarshelled2.LastName = "Simpson"
	structToBeMarshelled2.Age = 42

	mySlice = append(mySlice, structToBeMarshelled2)
	log.Println("Writing JSON from a Struct...")

	// In production we don't use MarshalIdent because this is just for visualization purposes
	newJson, error := json.MarshalIndent(mySlice, "", "   ")

	if error != nil {
		log.Println("Error while marshalling", error)
	}

	// We have to convert our bytes into a string so we can see the results:
	fmt.Println(string(newJson))

}

func readJSON() {
	jsonExample := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"age": 10
		},

		{
			"first_name": "Rivis",
			"last_name": "Poleselo",
			"age": 25
		}
	]
	
	`
	log.Println("Reading JSON...")
	// Since we don't know how many entries the JSON we have, we create a slice so it can populate the variable dynamically
	var unmarshelled []Person

	// In order to unmarshal we have to pass the bytes from the struct (as if we were processing a real JSON)
	err := json.Unmarshal([]byte(jsonExample), &unmarshelled)

	if err != nil {
		log.Println("Error unmarshelling JSON.", err)
	}

	// % v is for an interface
	log.Printf(("Unmarshalled %v"), unmarshelled)
}
