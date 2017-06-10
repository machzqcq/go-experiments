package main

import (
	"fmt"
)

func main() {
	//Variables to store course rankings

	firstRank := "39"
	secondRank := "614"

	if firstRank < secondRank {
		fmt.Println("\n First course is doing better " +
			"than second course")
	} else if firstRank > secondRank {
		fmt.Println("\n oh dear...your first" +
			"course must be doing abysmally")
	} else {
		fmt.Println("Both course are doing either" +
			"the same or something wierd is going on")
	}

	//Simple statement
	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("\n First course is doing better " +
			"than second course")
	}

}
