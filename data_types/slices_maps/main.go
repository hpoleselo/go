package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// A map is immutable and sorted, meaning we should always look for the values by the keys since they aren't in the same order!
	myStringMap := make(map[string]string)

	myStringMap["dog"] = "Robson"
	log.Println(myStringMap["dog"])

	myIntMap := make(map[string]int)
	myIntMap["first"] = 1
	myIntMap["second"] = 2
	log.Println(myIntMap["first"])

	// We can create a map of a struct
	userMap := make(map[string]User)
	me := User{
		FirstName: "Henrique",
		LastName:  "Poleselo",
	}

	userMap["me"] = me
	log.Println(userMap["me"])

	sliceExample()
}

func sliceExample() {
	// While maps are unsorted, arrays is a data structure that is sorted and can store many datatypes
	// Instead of using arrays we use slices in Go

	var mySlice []string
	mySlice = append(mySlice, "Henrivis")
	mySlice = append(mySlice, "Poleselo")
	log.Println(mySlice)

	var mySliceOfInts []int
	mySliceOfInts = append(mySliceOfInts, 2)
	mySliceOfInts = append(mySliceOfInts, 1)
	mySliceOfInts = append(mySliceOfInts, 3)
	sort.Ints(mySliceOfInts)
	log.Println(mySliceOfInts)

	// Assigning a slice with populated values
	mySlice2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	log.Println(mySlice2)
	log.Println(mySlice2[0:4])
}
