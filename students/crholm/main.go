package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		var buf = ""
		if i%3 == 0 {
			buf += "Fizz"
		}
		if i%5 == 0 {
			buf += "Buzz"
		}

		if buf != "" {
			fmt.Println(buf)
			continue
		}
		fmt.Println(i)
	}
}
