package main

var (
	message string = "Hello world" //use const for declaring constants
)

func main() {
	println(message)
}

//init runs after variable initialization but before any other function runs
func init() {
	message = "Hello Go"
}
