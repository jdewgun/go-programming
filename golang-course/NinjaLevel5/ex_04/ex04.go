package main

import (
	"fmt"
)

func main() {

	personOne := struct{
		firstName string
		lastName string
		programmingLanguages []string
	}{
		firstName: "John",
		lastName: "Doe",
		programmingLanguages: []string{"Python", "Golang", "JS"},
	}

	fmt.Printf("First Name of Person: %v \n", personOne.firstName)
	fmt.Printf("Last Name of Person: %v \n", personOne.lastName)
	fmt.Printf("Preferred Programming Languages: %v \n", personOne.programmingLanguages)
}