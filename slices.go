package main

import (
	"fmt"
)

func main() {

	//Declare a slice called myCourses
	myCourses := make([]string, 5, 10)
	fmt.Println("length is", len(myCourses), "and capacity is", cap(myCourses))

	//Delcare an int slice

	//mySlice := []int{1,2,3,4,5}
	//fmt.Println(mySlice[4])
	//sliceOfSlice := mySlice[2:5] this means 2,3,4 indexes
	//anothersliceOfSlice := mySlice[2:] , means index 2 through last index
}
