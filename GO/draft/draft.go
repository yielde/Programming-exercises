package main

import "fmt"

func main() {
	// fib(1): 1
	// fib(2): 1
	// fib(3): 2
	// fib(4): 3
	// fib(5): 5
	// fib(6): 8
	// fib(7): 13
	// fib(8): 21
	// fib(9): 34
	// fib(10): 55
	f := fib()
	for i := 1; i <= 10; i++ {
		if i <= 2 {
			fmt.Printf("fib(%d): %d\n", i, 1)
		} else {
			fmt.Printf("fib(%d): %d\n", i, f())
		}
	}
}

func fib() func() int {
	res, x := 1, 1
	return func() int {
		res, x = res+x, res
		return res
	}
}
