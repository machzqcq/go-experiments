package main

func main() {
	numTerms, result := add(1, 3, 5, 6, 7, 8)
	println("Added:", numTerms, "terms for a total of:", result)
}

func add(terms ...int) (numTerms int, result int) { //generally used for the pattern to return value and status code for the function execution
	for _, term := range terms {
		result += term
	}
	numTerms = len(terms)
	return
}
