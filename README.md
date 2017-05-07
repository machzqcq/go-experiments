# Golang experiments

Most folks switch to go for its concurrency model (from other programming languages). It is said Go has its own threading model aka. goroutines (comparable to green threads in python)
The reason I am trying out Go is because Docker was written in Go.  

Go has ~25 keywords

## Installation and initial setup

- Install Go (look up internet - pretty simple)
- Set up environment variables (GOPATH=<your code worspace> and GOROOT=<which go>)
- Example below
```
# On a mac
echo 'export GOPATH=$HOME/Code/go' >> ~/.bash_profile
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bash_profile
source ~/.bash_profile
mkdir -p ~/Code/go/src/github.com/machzqcq/go-experiments
```
- I set up my GOPATH=/Users/pmacharl/code/go and GOROOT=/usr/local/go (this was already set while installation)
- From shell/command line type `go env` - You should be able to see the environment variables

## IDE

- Installed Sublime Text 3
- Then install [package control](https://packagecontrol.io/installation)
- For Go intellisense - installed [GoSublime]() package from Preferences -> Browse Packages -> search for GoSublime and click install

# Concepts

## Variables, Constants, Pointers

- Have to use keyword 'var' declaring variables at package level. Global in scope
- Declare and initialize is done using ":=". This short assignment works only within functions and idiomatic way
- Declaring variables at the function level scope and not using causes compilation error
- "reflect" package used for checking type using `reflect.TypeOf(<value>)`
- implict inference of types is there for e.g. with "val" is automatically determined as string
- Explicit Type coercion - int(<stringvalue>) will convert string to integer
- Pointer value for a variable "x" is &x
- Dereferencing the value can be done by using * 
```
x :=2
ptr := &x
fmt.Println(Memory address of *x* variable is", ptr, "and the value of *x* is", *ptr)
```
- Constants can be done using `const` keyword
- Environment variables are gotten from "os" package
```
package main

import (
	"fmt"
	"os"
)

func main() {
	for _, env := range os.Environ() {
	fmt.Println(env)  //also see os.Getenv("USERNAME")
}
}
```


## Functions
- Functions pass parameters by value (though we can get around by passing the address of input parameters using &x, x being input parameter as in C pointers)
- Variadic functions can take variable number of input parameters
- General syntax
```
func funcName(value1 string) string {
	<code>
}
```
- '{' has to start on the same line as the function name

## If, switch

- if, else if and else (else is the last)
- switch case : break is implicit
- Explicit fallthrough - execute the next case block
- if in error handling (best practice in Go)
```
func testConn(target string) (rspTime float64 err error)

if err != nil {
		<error handling code>
}
<code>
```

## Loop

- Only one command/keyword for looping aka. for (this is the same keyword for infinite looop, while, looping over range/slice/map)
- Syntax
```
for <expression> {
	
}
for keyword
expression
	Blank expression = infinite loop
	Can be Boolean expression
	Can be range
Placement of curly braces is vital
```

## Arrays and Slices

- Slices are most likely used in Go (vs. Arrays). Slices are dynamic sized arrays. Slices are slices of arrays
- Slices are references. They are passed by reference. Slice data is always stored in an array
- Syntax : make(<type>, <len>, <cap>)
- append(<slice>, i) will generally double the capacity of the slice automatically, when you add an element i beyond the capacity of slice. This is used to expand the slice





