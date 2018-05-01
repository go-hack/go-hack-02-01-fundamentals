// Johan's FizzBuzz

package main

import "fmt"

func main() {
    for n := 1; n < 101; n += 1{
            fizzBuzz(n)
        }
}

func fizzBuzz(num int) {
    if num%3 == 0 && num%5 == 0 {
        fmt.Println( "FizzBuzz" )
    } else if num%3 == 0 {
        fmt.Println( "Fizz" )
    } else if num%5 == 0 {
        fmt.Println( "Buzz" )
    } else {
        fmt.Println( num )
    }
}
