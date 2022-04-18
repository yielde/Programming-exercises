/*
 * @Author: Galen Tong
 * @Date: 2022-04-18 23:46:59
 * @LastEditTime: 2022-04-18 23:53:12
 * @Description: 代码执行追踪
 */

package main

import "fmt"

func main() {
	// 	entering:  b
	// in b
	// entering:  a
	// in a
	// leaving:  a
	// leaving:  b
	b()
}

func trace(s string) string {
	fmt.Println("entering: ", s)
	return s
}

func untrace(s string) {
	fmt.Println("leaving: ", s)
}

func a() {
	defer untrace(trace("a"))
	fmt.Println("in a")
}
func b() {
	defer untrace(trace("b"))
	fmt.Println("in b")
	a()
}
