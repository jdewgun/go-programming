package main

import (
	"fmt"
)


func main() {
	string_literal := `This is a String Literal.

	String Literals are made using the backticks

	Now I understand why GCP BigQuery has backticks

	Or maybe I dont..`

	fmt.Println(string_literal)
}