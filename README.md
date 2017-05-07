# Golang experiments

Most folks switch to go for its concurrency model (from other programming languages). It is said Go has its own threading model aka. goroutines (comparable to green threads in python)
The reason I am trying out Go is because Docker was written in Go

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