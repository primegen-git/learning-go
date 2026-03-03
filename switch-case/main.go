package main

import "fmt"

func main() {
	fmt.Println("Learning switch case")

	fruits := []string{"banana", "apple", "grapes", "strawberry", "watermelon"}

	for _, fruit := range fruits {
		switch size := len(fruit); size { // scoping the size variable
		case 1, 2, 3, 4:
			fmt.Println(fruit, "is too short")
		case 5:
			WordLen := len(fruit) // scope of this variable is limited to this case block.
			fmt.Println(fruit, "is exactly of right length.", WordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(fruit, "is too long")

		}
	}
}
