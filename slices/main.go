package main

import "fmt"

func main() {
	fmt.Println("Learning slices")

	// different ways to declare a slices
	x := []int{10, 20, 30}
	var y []int

	fmt.Println(x)
	fmt.Println("nil slices:", y)
}
