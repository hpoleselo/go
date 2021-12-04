package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"simple_local_server/schemas"
	"simple_local_server/utils"
)

// Function that returns the division between two floats and if that operation fails because of impossible division by 0, an error is returned.
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero.")
		return 0, err
	}
	result := x / y
	return result, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	ps := &schemas.PayloadSchema{}
	err := utils.ValidateJSONBody(w, r, &ps)

	if err != nil {
		mr := &utils.MalformedRequest{}
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	fmt.Fprintf(w, "Person: %+v", ps)

	var x, y float32 = 100.0, 0.0
	f, err := divideValues(x, y)

	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0.")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("We're in the divide page and the division between %f and %f is: %f", x, y, f))
}
