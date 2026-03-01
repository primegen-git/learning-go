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

	// working with characters that is longer than one bytes
	var emoji_string string = "Good morning ☀️"
	fmt.Println("emoji_string: ", emoji_string)

	// len() method on string, return the length of the bytes not the length of the string characters.
	fmt.Println("emoji_string", len(emoji_string))

	// slices can cause problem with character longer than one byte.
	fmt.Println("slices of emoji_string: ", emoji_string[9:15]) // we have not taken all the characters of sun emoji.

	fmt.Println("slices of emoji_string: ", emoji_string[9:])

	// runes and bytes does not have one-to-one relationship
	var runes_emoji_message []rune = []rune(emoji_string)
	var bytes_emoji_message []byte = []byte(emoji_string)

	fmt.Println("runes_emoji_message: ", runes_emoji_message)
	fmt.Println("len(runes_emoji_message): ", len(runes_emoji_message))

	fmt.Println("bytes_emoji_message: ", bytes_emoji_message)
	fmt.Println("len(bytes_emoji_message): ", len(bytes_emoji_message))

}
