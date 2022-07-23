package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// fmt.Println(greetings.Hello("ok"))
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
