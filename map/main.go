package main

import "fmt"

func main() {
	fmt.Println("Leanring Map")

	// declare a nilMap
	var nilMap map[int]int
	fmt.Println("nilMap: ", nilMap)

	// declare a map
	num_map := map[string]int{
		"two":   2,
		"three": 3, // even the last element has the , at the end.
	}

	num_map["one"] = 1
	num_map["ten"] = 10

	fmt.Println("num_map: ", num_map)

	// create a map with make() method
	chain_map := make(map[string]float64, 10) // this length can be surpassed when chain_map is full.

	chain_map["ethereum"] = 4.7
	chain_map["solana"] = 4
	chain_map["sui"] = 4.2

	fmt.Println("chain_map: ", chain_map)

	// go return 0 if key does not exist in the map
	// ok Idiom (use to check whether the key exist in the map or not)

	m := map[string]int{
		"naruto":  5,
		"uzumaki": 0,
	}

	value, ok := m["naruto"]
	fmt.Printf("value: %d, ok: %t\n", value, ok)

	value, ok = m["uzumaki"]
	fmt.Printf("value: %d, ok: %t\n", value, ok)

	value, ok = m["sasuke"]
	fmt.Printf("value: %d, ok: %t\n", value, ok)

}
