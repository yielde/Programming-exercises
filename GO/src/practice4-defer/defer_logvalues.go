/*
 * @Author: Galen Tong
 * @Date: 2022-04-18 23:54:18
 * @LastEditTime: 2022-04-18 23:57:12
 * @Description:
 */
package main

import (
	"io"
	"log"
)

func main() {
	// 2022/04/18 23:57:24 func1("GO") = 1, EOF
	func3("GO")
}

func func3(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 1, io.EOF
}
