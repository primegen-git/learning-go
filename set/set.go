package main

import "fmt"

func main() {
	fmt.Println("Learning set")

	// define set with map
	intSet := map[int]bool{}

	var vals []int = []int{1, 2, 3, 2, 3, 4, 5, 6, 7}

	for _, val := range vals {
		intSet[val] = true
	}

	fmt.Println("intSet: ", intSet)
	fmt.Println("len(intSet): ", len(intSet))
	fmt.Println("len(vals): ", len(vals))

	fmt.Println("Noticed that len(vals) != len(intSet), which means the duplicate values from vals are ignored.")
}
