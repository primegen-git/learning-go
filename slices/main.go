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

	fmt.Println("nz: ", nz)

	// observering the change in capacity as len increases.

	var nxx []int
	fmt.Printf("len of nxx: %d, cap of nxx: %d\n", len(nxx), cap(nxx))

	nxx = append(nxx, 10)
	fmt.Printf("len of nxx: %d, cap of nxx: %d\n", len(nxx), cap(nxx))

	nxx = append(nxx, 20)
	fmt.Printf("len of nxx: %d, cap of nxx: %d\n", len(nxx), cap(nxx))

	nxx = append(nxx, 30)
	fmt.Printf("len of nxx: %d, cap of nxx: %d\n", len(nxx), cap(nxx))

	nxx = append(nxx, 40)
	fmt.Printf("len of nxx: %d, cap of nxx: %d\n", len(nxx), cap(nxx))

	nxx = append(nxx, 50)
	fmt.Printf("len of nxx: %d, cap of nxx: %d\n", len(nxx), cap(nxx))

	// slicing in slices

	nyy := []int{10, 20, 30}
	nzz := nyy[1:]

	fmt.Println("nyy: ", nyy)
	fmt.Println("nzz: ", nzz)

	// showing that slicing point to same memory location
	nzz[0] = 15

	fmt.Println("nyy: ", nyy)
	fmt.Println("nzz: ", nzz)

}
