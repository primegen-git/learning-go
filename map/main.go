package main

import "fmt"

func main() {
	fmt.Println("Leanring Map")

	// declare a nilMap
	var nilMap map[int]int
	fmt.Println("nilMap: ", nilMap)

	// declare a map
	num_map := map[string]int{}

	num_map["one"] = 1
	num_map["ten"] = 10

	fmt.Println("num_map: ", num_map)

}
