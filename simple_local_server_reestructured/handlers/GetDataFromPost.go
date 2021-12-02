package handlers

import (
	"log"
	"net/http"
)

// handlers is the name of the package (folder)
// And each handler is a file

// TODO: Add validate field
// field string `json:"modelGUID" validate:"required"`
type PayloadSchema struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func ParseDataFromPost(w http.ResponseWriter, r *http.Request) {
	log.Println("Received POST request, parsing data...")
	// Set Response to be a JSON
	w.Header().Set("Content-Type", "application/json")
}
