package main

import "fmt"

func main() {
	fmt.Println("Leanring struct")

	// declaring a struct
	type person struct {
		name string
		age  int
	}

	// creating variable for struct

	var sasuke person
	fmt.Println("sasuke: ", sasuke)

	// assigning values without using field of struct

	var naruto person = person{"uzumaki naruto", 10}
	fmt.Println("naruto: ", naruto)

	// assigning with using the field of struct
	var rock_lee person = person{
		name: "rock_lee",
		age:  30,
	}
	fmt.Println("rock_lee: ", rock_lee)

	// anonymous structs

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println("anonymous struct: ", pet)

}
