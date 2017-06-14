package main

import (
	"fmt"
)

func main() {
	mySlice := make([]float32, 100)
	mySlice[0] = 1.
	mySlice[1] = 2.
	mySlice[4] = 12.
	fmt.Println(mySlice)
}
