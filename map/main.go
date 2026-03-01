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

	// create a map with make() method
	chain_map := make(map[string]float64, 10) // this length can be surpassed when chain_map is full.

	chain_map["ethereum"] = 4.7
	chain_map["solana"] = 4
	chain_map["sui"] = 4.2

	fmt.Println("chain_map: ", chain_map)

}
