package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Get a gretting message and print it.
	log.SetPrefix("grettings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Gladys")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(message)
}
