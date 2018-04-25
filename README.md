# 02-01 Fundamentals

This section is to get a brief overview of go and our development 
environment, mainly Visual Studio Code


## Visual Studio Code
This editor is by default a relative small editor that are made by Microsoft using Github Electron framework. It is liked by many for 
development in Go, which is the language we will be using. VS Code will help us in writing code, highlight errors, auto compleat things and so on. While you can write code in any text editor, a good on can be of great help, especially in the beginning.

### Extensions
VS Code comes quite bare, but can be augmented with extensions customizing your experience. For programming in go i recommend that you install the `Go` extensions by `lukehoban`.

So lets install some extensions.. Install extensions by pressing `Ctrl+Shift+X` and search for `Go`, once find just click install, and your done.


## Lets learn Go

### Compiler
For now we will keep things simple and not go through everything that go compiler and tooling around it can do, this will come when we get into it later. The easiest way to interact and compile go software is through a terminal. In Linux or Mac you can simply open a terminal and navigate to the directory where you source code is located. In windows, you instead use the more clumsy `command prompt` but the command should be the same

Running your program is done by the command `go run <your file>`
e.g.
```bash
$ go run main.go
``` 

### Code
You can always find the language specification for go on
[https://golang.org/ref/spec](https://golang.org/ref/spec)
and is god to have as a reference, but [https://golang.org/doc/effective_go.html](https://golang.org/doc/effective_go.html) is usually better if you want to know how to use it.

All go programs starts with in the same way and we have seen it in the hello-world example, so lets go through it with some explanations

```go

// +- This tells the go compiler that this file is in the packet main.
// |  The main package is special and is expected to contain a
// |  function called `main` which is the entry point to your program.
// |   
// v
package main    


// +- import statements tells go which other packets are being used 
// |  in this file and where to look for code. When you program you
// |  dont have to reinvent the wheal every time and can use code 
// |  others have written
// |   
// v
import "fmt"

// +- This is a function called main. The word `func` tells go that 
// |  you are describing a function and the word `main` names this  
// |  function in further use. The function called main is special 
// |  and is the entry point when you run your program. This is where 
// |  the execution starts
// |
// v
func main() {
    // +- This calls the function Println in the packet fmt and
    // |  passes a string "Hello world!" as a parameter. 
    // |  The function will print the parameter passed to the 
    // |  screen and appending a new line.
    // v
    fmt.Println("Hello world!")
}

```





## Lets do some coding

### Cloning and starting work

1. Find your go path, a dir named `go` usually located in your home directory. 
1. Navigate and make the following folder structure `$GOPATH/src/github.com/<YOUR GITHUB NAME>`
1. Now run `git clone https://github.com/<YOUR GITHUB NAME>/go-hack-02-01-fundamentals.git`, this will fetch this 
   project for you to work on 
1. Open Visual Studio Code and select the folder `$GOPATH/src/github.com/<YOUR GITHUB NAME>/go-hack-02-01-fundamentals`
   to work from 

I now assume you have visual studio code open aswell as the project within it. 

1. Let us not begin by creating a folder `students/<YOUR GITHUB NAME>`
1. Now create an empty file within the folder called `main.go`