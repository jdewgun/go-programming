package main

import (
	"fmt"
	"encoding/json"
)

type user struct {
	First string
	Age   int
}

func main() {
	user1 := user{
		First: "James",
		Age:   32,
	}

	user2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	user3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{user1, user2, user3}

	fmt.Println(users)

	// your code goes here

	data, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
