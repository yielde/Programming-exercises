/*
 * @Author: Galen Tong
 * @Date: 2022-04-17 14:47:02
 * @LastEditTime: 2022-04-21 10:26:07
 * @Description:
 */
package main

import (
	"fmt"
	"time"
)

const LIM = 41

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
	// fib(11): 89
	// fib(12): 144
	// fib(13): 233
	// fib(14): 377
	// fib(15): 610
	// fib(16): 987
	// fib(17): 1597
	// fib(18): 2584
	// fib(19): 4181
	// fib(20): 6765
	// fib(21): 10946
	// fib(22): 17711
	// fib(23): 28657
	// fib(24): 46368
	// fib(25): 75025
	// fib(26): 121393
	// fib(27): 196418
	// fib(28): 317811
	// fib(29): 514229
	// fib(30): 832040
	// fib(31): 1346269
	// fib(32): 2178309
	// fib(33): 3524578
	// fib(34): 5702887
	// fib(35): 9227465
	// fib(36): 14930352
	// fib(37): 24157817
	// fib(38): 39088169
	// fib(39): 63245986
	// fib(40): 102334155
	// fib(41): 165580141
	// used time:  2.338626802s
	start := time.Now()
	for i := 0; i < LIM; i++ {
		res := fib(i)
		if i <= 1 {
			fmt.Printf("fib(%d): %d\n", i+1, res)
		} else {
			fmt.Printf("fib(%d): %d\n", i+1, res)
		}
	}
	end := time.Now()
	interval := end.Sub(start)
	fmt.Println("used time: ", interval)
}

func fib(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fib(n-1) + fib(n-2)
	}
	return
}
