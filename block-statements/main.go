package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Learn Block statements")

	// if statement

	// way 1:
	n := rand.Intn(10)

	if n < 5 {
		fmt.Println("n is low.")
	} else if n > 5 {
		fmt.Println("n is high.")
	} else {
		fmt.Println("n is perfect.")
	}

	// way 2:

	if n := rand.Intn(10); n < 5 { // the scope of n in here is limited if blocks.
		fmt.Println("n is low.")
	} else if n > 5 {
		fmt.Println("n is high.")
	} else {
		fmt.Println("n is perfect.")
	}

	// for loop (classic one)
	for i := 1; i < 10; i++ {
		fmt.Println("i: ", i)

	}

	// conditional-only for loop  (this is like while loop for other langauge)
	i := 1

	for i < 10 {
		fmt.Println("i: ", i)
		i = i + 1
	}

	// infinite loop

	// for {
	// 	fmt.Println("Infinite loop, pres Ctrl+c")
	// }

	// for - range loop

	evenVals := []int{2, 4, 6, 8, 10, 12}

	for i, v := range evenVals {
		fmt.Printf("index: %d, val : %d\n", i, v)
	}

	// iterating over a map
	numbers := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	// notices: that the value printed by the map is different each time, map does not have ordered iternation
	for j := range 3 {
		fmt.Println("loop - ", j)
		for k, v := range numbers {
			fmt.Printf("key: %d, value: %s\n", k, v)
		}

	}

	// iteration over string
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

}
