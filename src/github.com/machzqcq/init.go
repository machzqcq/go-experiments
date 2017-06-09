package main

import (
	"fmt"
)

var (
	message string = "Hello world" //use const for declaring constants
)

func main() {
	fmt.Println(message)
}

//init runs after variable initialization but before any other function runs
func init() {
	message = "Hello Go"
}
