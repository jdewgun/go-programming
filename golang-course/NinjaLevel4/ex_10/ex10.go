package main

import (
	"fmt"
)

func main() {
	dictionary := map[string][]string {
		`JamesBond`: []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`MoneyPenny`: []string{`James Bond`, `Literature`, `Computer Science`},
		`DrNo`: []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println("Map before delete", dictionary)

	delete(dictionary, `DrNo`)

	fmt.Println("Map after delete", dictionary)

	for key, value := range dictionary {
		fmt.Printf("Key: %v \t Value: %v \n", key, value)
		for i, v := range value {
			fmt.Println("\t", i, v)
		}
	}
}