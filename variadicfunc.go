package main

import (
	"fmt"
)

func main() {
	bestFinish := bestLeagueFinishes(13, 10, 11, 9, 8, 7, 1, 2)
	fmt.Println(bestFinish)
}

func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
