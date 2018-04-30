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


#### Primitives and variables

There are, as in most programming languages some primitive types.
`bool` a true or false value, `string` a piece of text, `int` a integer number, `float` a decimal number, `byte` a int with 8bits, `run` a character. There are others but lets stay with this for now.  These are used to construct everything else. An important thing about primitives is that they always have a value, even if you dont assign one to them.


A variable is can be views as simply naming a value with an alias. That value of the alias can be used in conjunction with others. The slightly trick part with go is that there are two ways to declare a variable. 

```go
// this declares a variable, myVariable1, of type string and assigns it
// the defaul value of an empty string
var myVariable1 string

// this declares a variable, myVariable2, of type string and assigns it
// the value World
myVariable2 := "World"

// this reassigns a variable, myVariable1, and assigns it
// to the value hello
myVariable1 = "Hello"

// this reassigns a variable, myVariable1, as the concatination result of
// "Hello" + "World" == "HelloWorld"
myVariable1 = myVariable1 + myVariable2
```

#### Control structure 

`if` statements are used to control your flow of the program. If somthing is true, do something. 
```go

// this declares an int and assigns it its default value 0
var a int

if a == 0 {
    fmt.Println("Its Zero")
}

```


`for` is the keyword for a loop, this is used if you want to do something many times
```go

// this declares an int and assigns it its default value 0
var sum int

//     +- This delcares a in variable i as 0
//     |      +- As long as this is true the loop will continue.
//     |      |        +- When each loop is done, i is incremented by 1
//     V      V        V
for i := 0; i < 10; i += 1{
    sum += i+1
}

// The prints sum of every number between 1 to 10
fmt.Println(sum)

```

## Lets do some coding

### Cloning and starting work

1. Find your go path, a dir named `go` usually located in your home directory. 
1. Navigate and make the following folder structure `$GOPATH/src/github.com/<YOUR GITHUB NAME>`
1. Now run `git clone https://github.com/<YOUR GITHUB NAME>/go-hack-02-01-fundamentals.git`, this will fetch this 
   project for you to work on 
1. Open Visual Studio Code and select the folder `$GOPATH/src/github.com/<YOUR GITHUB NAME>/go-hack-02-01-fundamentals`
   to work from 

I now assume you have visual studio code open as well as the project within it. 

1. Let us not begin by creating a folder `students/<YOUR GITHUB NAME>`
1. Now create an empty file within the folder called `main.go`

### Fizz Buzz 
The assignment is to write a program simulates the game FizzBuzz. It should print every number between 1 and 100. But 

* If the number is evenly divisible by 3, replace the number by the word Fizz.
* If the number is divisible by 5, replace the number by the word Buzz. 
* If the number is divisible by 3 and 5, replace the number by the word FizzBuzz

Expected output

```bash
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
.
.
.
```

A good thing to know here integer operators. 
* `+` adds two numbers, `5 + 2 = 7` 
* `-` subtract one number form another, `5 - 2 = 3` 
* `*` multiplies a number by another, `5 * 2 = 10`
* `/` divides a number by another without remainder, `5 / 2 = 2`
* `%` calculates the remainder in a division, `5 % 2 = 1`

Some notes here on integer operations. In an integer operation 
`(5 / 2) * 2 + (5 % 2) = 5`. Meaning, by dividing by 2 and multiplying the result by 2 dose not yield what we started with, five in this case, since we loose information in the precision of integers.
