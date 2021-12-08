package main

import (
	"fmt"
	"net/http"
	"simple_local_server/handlers"
)

// cannot be changed in the application
const portNumber = ":8080"

func main() {

	http.HandleFunc("/about", handlers.About)

	http.HandleFunc("/divide", handlers.Divide)

	http.HandleFunc("/parse_data", handlers.ParseDataFromPost)

	fmt.Println(fmt.Sprintf("Starting application on port %s...", portNumber))
	// Start a web server to listen for requests, ignoring error messages btw
	_ = http.ListenAndServe(portNumber, nil)
}
