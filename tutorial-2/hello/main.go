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

	names := []string{"rock lee", "naruto", "sasuke", "guy"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for name, message := range messages {
		fmt.Printf("%s: %s\n", name, message)
	}
}
