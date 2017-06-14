package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}

	i := 0
	for { //similar to while loop, when you don't know when the loop should stop at compile time
		i++
		println(i)

		if i > 5 {
			break
		}
	}
}
