package main

import (
	"fmt"
)

func main() {

	coursesInProg := []string{"Docker deep dive",
		"Docker Clustering", "Docker and Kubernetes"}

	coursesCompleted := []string{"Docker deep dive",
		"go fundamentals", "Puppet fundamentals"}

	for _, i := range coursesInProg {
		//fmt.Println(i)
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("\n Hey we found a clash.",
					i, "is in both lists")
			}
		}
	}
}
