package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib1, fib2 := 1, 0
	return func() int {
		fib2, fib1 = fib1, fib1 + fib2
		return fib1 - fib2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
