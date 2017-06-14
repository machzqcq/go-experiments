package main

import "fmt"

func main() {
	myMap := make(map[int]string) //string values organized by integers

	fmt.Println(myMap)

	myMap[3] = "blah"
	myMap[1000] = "xyz"

	fmt.Println(myMap)
	fmt.Println(myMap[45])
}
