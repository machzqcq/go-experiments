package main

func main() {
	emp := enhancedMessagePrinter{} //enhancedMessagePrinter{messagePrinter{"foo"}} works, but not best practice
	emp.message = "foo"
	emp.printMessage()

}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}

type enhancedMessagePrinter struct {
	messagePrinter
}
