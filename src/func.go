package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "function basics"
	author := "nigel poulton"

	fmt.Println(converter(module, author))
}

func converter(module, author string) (moduleToTitle, authorToUpper string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}
