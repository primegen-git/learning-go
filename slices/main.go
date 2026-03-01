package main

import "fmt"

func main() {
	fmt.Println("Learning slices")

	// different ways to declare a slices
	x := []int{10, 20, 30}
	var y []int

	fmt.Println(x)
	fmt.Println("nil slices:", y)

	// len built-in command
	fmt.Println("Length of x: ", len(x))
	fmt.Println("Length of y: ", len(y))

	// append to new array
	var nx []int

	nx = append(nx, 10)

	fmt.Println("nx: ", nx)

	// append to existing array
	var ny = []int{10, 20, 30}

	ny = append(ny, 40, 50)

	fmt.Println("ny: ", ny)

	// append to an existing slices

	var nz = []int{5}

	nz = append(nz, ny...) // ... is like expanding the ny as append(nz, 10, 20, 30, 40, 50)

}
