package main

func main() {
	message := "hello"
	sayHello(&message) //& gives memory address
	println(message)
}

func sayHello(message_mem_address *string) { //pointer i.e. memory address
	println(*message_mem_address) //* will get the value i.e. derefencing a pointer

	*message_mem_address = "hello go"
}
