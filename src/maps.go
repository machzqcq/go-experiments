package main

import (
	"fmt"
)

func main() {
	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}

	for key, value := range testMap {
		fmt.Printf("\n key is: %v Value is: %v", key, value)
	}

	testMap["E"] = 5
	fmt.Println(testMap)

	//Delete
	// delete(testMap, "A")
}
