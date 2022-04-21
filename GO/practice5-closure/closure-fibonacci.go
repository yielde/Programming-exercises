/*
 * @Author: Galen Tong
 * @Date: 2022-04-20 18:11:53
 * @LastEditTime: 2022-04-20 18:39:08
 * @Description: closure fibonacci and closure debug
 */
package main

import (
	"fmt"
	"log"
	"runtime"
)

var where = func() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}

func main() {
	// 2022/04/20 18:39:11 /home/galen/Programming-exercises/GO/practice5-closure/closure-fibonacci.go:32
	// 2022/04/20 18:39:11 /home/galen/Programming-exercises/GO/practice5-closure/closure-fibonacci.go:40
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
	where()
	for i := 1; i <= 10; i++ {
		if i <= 2 {
			fmt.Printf("fib(%d): %d\n", i, 1)
		} else {
			fmt.Printf("fib(%d): %d\n", i, f())
		}
	}
	where()
}

func fib() func() int {
	res, x := 1, 1
	return func() int {
		res, x = res+x, res
		return res
	}
}
