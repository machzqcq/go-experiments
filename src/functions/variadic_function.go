package main

func main() {
	sayHello("Hello", "Go")
}

func sayHello(message ...string) {
	for _, message := range message {
		println(message)
	}
}
