package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"simple_local_server/schemas"
	"simple_local_server/utils"
)

// Source: https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

// handlers is the name of the package (folder)
// And each handler is a file

// TODO: Add validate field
// field string `json:"modelGUID" validate:"required"`
//https://stackoverflow.com/questions/62522747/golang-validator-with-custom-structs

func ParseDataFromPost(w http.ResponseWriter, r *http.Request) {
	pschema := &schemas.PayloadSchema{}

	log.Println("Received POST request, parsing data...")

	err := utils.ValidateJSONBody(w, r, pschema)

	if err != nil {
		log.Println("JSON Body isn't valid: ", err)
	}

	// Set Response to be a JSON
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	// Marshall data to return as JSON as well

	marshelledData, err := json.Marshal(pschema)
	if err != nil {
		log.Println("Could not marshal retrieved payload.")
		http.Error(w, "Could not marshal retrieved payload.", http.StatusBadRequest)
	}

	w.Write(marshelledData)
	log.Println("Data has been parsed and returned.")
}
