# Golang experiments

Most folks switch to go for its concurrency model (from other programming languages). It is said Go has its own threading model aka. goroutines (comparable to green threads in python)
The reason I am trying out Go is because Docker was written in Go.  

Go has ~25 keywords. Go focuses on Simplicity and Clarity

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

## Workspaces
As you continue working - you will quickly realize that you need multiple workspaces/projects i.e. you would not want to put your source code always in the Global GOPATH ($HOME/Code/go) as I declared above. There are two ways to deal with it

- Option1: By default you can simply have multiple paths defined in your GOPATH and Go looks for src in all folders in that order (So lets say you can have your global , private, work etc.). 
- Option2: Everytime your cd into the root of your projet (think of github project for e.g.), then initialize GOPATH. So essentially override the 'cd' command for bash shell and drop in a .gopath empty file in the root of every folder/project. That way the cd will append your current root of folder to GOTPATH. I had this in my ~/.bash_profile and a .gopath in the root of every go project where my 'src' resides OR `source ~/.bash_profile` if you had the below in your bash profile

```
cd () {
    builtin cd "$@"
    cdir=$PWD
    while [ "$cdir" != "/" ]; do
        if [ -e "$cdir/.gopath" ]; then
            export GOPATH=$cdir # GOPATH=$GOPATH:$cdir if you want to preserve existing
            break
        fi
        cdir=$(dirname "$cdir")
    done
}
```  

Note: If you are running `go build`, then run it from the project root and not from 'src' folder. So it might seem a little counterintuitive, but go build automatically looks in the $GOPATH/src folder i.e. you don't have to specifically step inside 'src' folder. This might be more evident in nested package structures. For e.g. if you had $GOPATH=~/git-projects/go-web/src/github.com/machzqcq/fileserver/fileserver.go, then when you cd into ~/git-projects/go-web (assuming you have a .gopath here) , then to build the binary, just run `go build github.com/machzqcq/fileserver` , this will put the fileserver binary in `~/git-projects/go-web/` folder. On the other hand `go install github.com/machzqcq/fileserver` drop the binary in  `~/git-projects/go-web/` folder

Note: If you are using vscode, then your integrated terminal executes .bashrc, hence include the above in .bashrc as well

- Related info: This might be the most closest to many developers who know RVM, pyvm, virtualenv etc. Here is the [gvm](https://github.com/moovweb/gvm) (go version manager). This lets you define multiple GOROOT (i.e. multiple Go runtimes) and manage multiple GOPATHs with gvm pkgset. Read the full documentation

## IDE

- Installed Sublime Text 3
- Then install [package control](https://packagecontrol.io/installation)
- For Go intellisense - installed [GoSublime]() package from Preferences -> Browse Packages -> search for GoSublime and click install

Alternately, you can use Visual Studio Code  

- Install VS Code from internet
- Extensions -> Install Go plugin from lukehoban - Thats it!

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
- Referencing a slice by its variable name references the entire slice
- for range looop iterate slices

## Maps
- Similar to slices and arrays in that they are lists
- Maps are UNORDERED lists. They are <key,value> pairs. Maps are references
- Both keytype and valuetype must be defined upfront since Go is strongly type
- Declaration Syntax : var myMap := make(map[keyType]valueType) 
```
myMap := map[string]int {
	"string1": 1,
	"string2": 2
}

for key, value := range myMap {
	fmt.Printf("\n key is", key, "value is", value)
}
```

- Maps are unsafe for concurrency
- Cheap to pass around
- Specify size for large maps would improve performance

## Structs

- Define custom data types
- Go is not designed for Object Oriented programming (NO objects, classes and inheritance)
- 'type' is the keyword for struct

## WebServer and FileServer
For managing multiple workspaces aka. GOPATH, see at the top of this page 

- From the root of this project, execute `go install github.com/machzqcq/firstapp` for http server and `go install fileserver` for file server. This will create a `bin/firstapp` folder
- Note that `go install` or `go build` work at package level ONLY, however `go run` can work at file level. So `go install github.com/machzqcq/firstapp/main.go` is NOT allowed
- The binaries get created in `go/bin` folder
- For file server, static html files were located in src/public folder
- So running `nohup bin/fileserver > nohup.out &` or `nohup bin/fileserver > /dev/null 2>&1 &` should run the program in background. The first one outputs to nohup.out file, second one shoves everything into pretty much a blackhole

### Handlers

- http.Handle requires to implement the ServeHTTP interface. When http.Handle is called, the method ServeHTTP is registered (Actually the nil means the Default behavior is to register ServeHTTP method)
```
package main

import (
	"net/http"
)
type myHandler struct {
	greeting string
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}

func main() {
	http.Handle("/", &myHandler{greeting: "Hello"})
	http.ListenAndServe(":8000",nil)
}
``` 
- Most of the times, http.HandleFunc() is good enough. See firstapp package
- Built-in handlers: NotFoundHandler,RedirectHandler,StripPrefix,TimeoutHandler,FileServer
- http.NotFoundHandler : Returns status 404, not found status to requestor. Administrative area in your application that requires user access. You could just forward them to this handler if they are not authorized to access , instead of writing up a whole new html page 
- http.RedirectHandler: Takes the request that came in and redirect to another url. Temporary redirect or permanent redirect
- http.StripPrefix: Looks kind of middleware. `func StripPrefix(prefix string, h handler) Handler`. Decorates another handler. Basically lets you handle a request from a url that has a prefix on it, but redirect to a handler that is not expecting the prefix
- http.TimeoutHandler: Again a decorator. `func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler` - First parameter is the handler that is decorating. The second argument is the amount of time that the first handler is allowed to process. IF timeout, then third argument is returned
- http.FileServer - `func FileServer(root FileSystem) Handler`. FileSystem is an interface, which is generally the interface to the OS filesystem. So for e.g. `type Dir string` implements the FileSystem interface. Most commonly we use for files on disk, hence Dir is used. The string value you pass becomes the root or current folder



