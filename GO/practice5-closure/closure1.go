/*
 * @Author: Galen Tong
 * @Date: 2022-04-20 15:50:00
 * @LastEditTime: 2022-04-20 15:54:15
 * @Description:
 */
package main

import "fmt"

func main() {
	// Adder(1): 1
	// Adder(2): 3
	// Adder(3): 6
	a := Adder()

	fmt.Printf("Adder(1): %d\n", a(1))
	fmt.Printf("Adder(2): %d\n", a(2))
	fmt.Printf("Adder(3): %d\n", a(3))
}

func Adder() func(x int) int {
	var y int
	return func(x int) int {
		y += x
		return y
	}
}
