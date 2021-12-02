package handlers

import (
	"fmt"
	"net/http"
)

// ----- Util functions -------

func addValues(x, y int) int {
	return x + y
}

// This function represents our home.
// In order to a function to respond to a Web Browser, it has to have 2 params, the writerResponse (w) and the Request (r)
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About page and we're calling a function inside it: %d ", sum))
}
