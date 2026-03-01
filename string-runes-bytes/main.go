package main

import "fmt"

func main() {
	fmt.Println("string, runes, bytes")

	// declare a string
	var welcome_msg string = "hello, world"
	fmt.Println("welcome message: ", welcome_msg)

	// access a byte using index
	var second_character byte = welcome_msg[1]
	fmt.Println("second_character: ", second_character) // this would be a number, because strings are stored in bytes and e = 101 (ASCII)

	// rune
	var s rune = 's'
	fmt.Println("rune s: ", s)

	// converting strings into a runes of character
	var welcome_msg_runes []rune = []rune(welcome_msg)
	fmt.Println("welcome_msg_runes: ", welcome_msg_runes) // go return a slice of bytes, go does not represent array of string as unicode.

	// to get the string, you have to use
	fmt.Println("second welcome_msg character: ", string(welcome_msg_runes[1]))
}
