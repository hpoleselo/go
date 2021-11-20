package main

import (
	"errors"
	"fmt"
	"net/http"
)

// ----- Util functions -------

func addValues(x, y int) int {
	return x + y
}

// Function that returns the division between two floats and if that operation fails because of impossible division by 0, an error is returned.
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero.")
		return 0, err
	}
	result := x / y
	return result, nil
}

// ----- Routes -----

func Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float32 = 100.0, 0.0
	f, err := divideValues(x, y)

	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0.")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("We're in the divide page and the division between %f and %f is: %f", x, y, f))
}

// This function represents our home.
// In order to a function to respond to a Web Browser, it has to have 2 params, the writerResponse (w) and the Request (r)
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About page and we're calling a function inside it: %d ", sum))
}
