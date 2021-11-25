package main

import "log"

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

// Now we want to add actions to our structs, so we create a generic interface
type Animal interface {
	// We list everything an animal can DO, basically we're listing the functions/methods and what they return
	Says() string
	HasNumberOfLegs() int
}

// Now we make a function to print the Animal's information by passing the wished animal (struct)
func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.HasNumberOfLegs(), "legs")
}


func (d Dog) Says() string {
	return "Woof"
}

func (d Dog) HasNumberOfLegs int {
	return 4
}

func main() {
	dog := Dog{
		Name:  "Simpson",
		Breed: "Pug",
	}

	gorilla := Gorilla{
		Name:          "Margaret",
		Color:         "Brown",
		NumberOfTeeth: 12,
	}

	// Now we try to print the information about the animal
	// But passing only dog as parameter doesn't work as dog doesn't implement the necessary interface (methods), that's why we have to declare a function just for that
	PrintInfo(dog)

}
