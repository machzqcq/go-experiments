package main

func main() {
	result := add(1, 3, 5, 6, 7, 8)
	println(result)
}

func add(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result
}
