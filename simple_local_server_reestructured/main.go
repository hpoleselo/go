package main

import (
	"fmt"
	"net/http"
)

// cannot be changed in the application
const portNumber = ":8080"

func main() {

	// When user hits /, we redirect user a response by passing the function in the second argument
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// We use Fprintf to print the response, that's why we pass w

		// FprintF returns an int (bytes) and an error (string)
		n, err := fmt.Fprintf(w, "Hello World")

		// If there is an error we print it out
		if err != nil {
			fmt.Println(err)
		}

		// Using Sprintf to use any type and present as a string
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))

	})

	http.HandleFunc("/about", About)

	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s...", portNumber))
	// Start a web server to listen for requests, ignoring error messages btw
	_ = http.ListenAndServe(portNumber, nil)
}
