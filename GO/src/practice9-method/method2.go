/*
 * @Author: Galen Tong
 * @Date: 2022-04-23 00:11:06
 * @LastEditTime: 2022-04-23 00:13:11
 * @Description: 数组方法
 */
package main

import "fmt"

type IntVector []int

func main() {
	// 6
	fmt.Println(IntVector{1, 2, 3}.Sum())
}

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}
