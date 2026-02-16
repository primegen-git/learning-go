package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a gretting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Print(message)
}
