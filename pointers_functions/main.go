// package decoration
package main

// format package
import "fmt"

// package level variable, available to anything inside this main package
var anyVariable string

// One go program has to have only one main and it must not return anything and not take any paremeter
func main() {
	fmt.Println("Hello World")

	var whatToSay string

	// int by default is a 64 length int since we're using a 64 bit comptuer
	var i int

	whatToSay = "I Said What About It"
	fmt.Println(whatToSay)
	i = 7
	fmt.Println("i is set to:", i)

	// The double dots assigns the type of whatever was returned by the saySomething function
	// In this case whatWasSaid will be a string
	whatWasSaid, otherThingThatWasSaid := saySomething()
	fmt.Println("The functions saySomething() returned:", whatWasSaid, otherThingThatWasSaid)

	// Pointer is a reference to a variable, location in memory
	var myString string
	myString = "Green"
	fmt.Println("myString is set to:", myString)

	// In order to change myString to Red by using the function we have to pass the reference of that variable, not the variable itself
	changeUsingPointer(&myString)

	// We separate the scope of variables here by not changing the scope of the variable
	fmt.Println("myString is set to:", myString)
}

// This function returns two strings
func saySomething() (string, string) {
	return "Something", "Something Else"
}

// This function takes "s" as argument of type string with pointer
// It does not return anything as we're not changing the scope of the variable, but we're referencing it in the main
//
func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
