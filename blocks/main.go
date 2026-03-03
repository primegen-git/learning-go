package main

import "fmt"

func main() {
	fmt.Println("Learning blocks")

	// understanding scoping

	x := 10

	if x > 5 {

		fmt.Println("start of if block")
		fmt.Println("x: ", x)

		fmt.Println("redeclared x")
		x := 5
		fmt.Println("x: ", x)
		fmt.Println("end of if block")
	}

	fmt.Println("x: ", x)

}
