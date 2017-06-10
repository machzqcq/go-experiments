// see selinux.go file and see this pattern
package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("/Users/pmacharl/blah.txt")

	if err != nil {
		fmt.Println("Error returned was :", err)
	}
}
