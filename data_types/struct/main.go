package main

import "log"

type myStruct struct {
	// We make it accessible to other packages by using capital letters on the attributes
	FirstName string
}

// We define a function to a struct by simply declaring a function
// Note that the function doesn't receive any argument, but it is preeceded by a receiver (m *myStruct) and returns a string
// The receiver ties the function printFirstName with the struct, that's why we use the pointer *myStruct and we have to pass a variable to it, which is m
func (m *myStruct) printFirstName() string {
	// So to access the struct we just have to reference the variable!
	return m.FirstName
}

func main() {

	// Two ways of actually instantiating the struct to a variable:
	var myVar myStruct
	myVar.FirstName = "Henry"

	// Or we could assign in an easier way
	myVar2 := myStruct{
		FirstName: "Poleselo",
	}

	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar2 is set to", myVar2.FirstName)

	// So now we can do the samething as above but using the structs's function:
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar is set to", myVar2.printFirstName())
}
