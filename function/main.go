package main

import "fmt"

func div(numerator float64, denominator float64) float64 {
	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}

// simulating optional parameter

type MyFuncOpts struct {
	FirstName string
	LastName  string
	age       any
}

func MyFunc(opts MyFuncOpts) {
	// do nothing
	fmt.Println("optional Parameter: ", opts)
}

func main() {
	fmt.Println("Learning function")

	fmt.Println("MyFuncOpts Struct: ", MyFuncOpts{"FirstName", "SecondName", "age"})

	MyFunc(MyFuncOpts{
		LastName: "naruto",
		age:      20,
	})
	MyFunc(MyFuncOpts{
		FirstName: "uzumaki",
		age:       20,
	})

}
