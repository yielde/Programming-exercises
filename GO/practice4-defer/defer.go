/*
 * @Author: Galen Tong
 * @Date: 2022-04-18 23:32:21
 * @LastEditTime: 2022-04-18 23:45:27
 * @Description:
 */
package main

import "fmt"

func main() {
	// 不加defer
	// func1 top
	// func2: deferred until the end of the calling function
	// func1 bottom
	// func1 finished

	// 加defer
	// func1 top
	// func1 bottom
	// func2: deferred until the end of the calling function
	// func1 finished
	func1()
	fmt.Println("func1 finished")

	// 多个defer被注册，逆序执行（栈）
	// 5
	// 4
	// 3
	// 2
	// 1
	stack()
}

func func1() {
	fmt.Println("func1 top")
	defer func2()
	fmt.Println("func1 bottom")
}

func func2() {
	fmt.Println("func2: deferred until the end of the calling function")
}

func stack() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}

