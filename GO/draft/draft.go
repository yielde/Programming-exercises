/*
 * @Author: Galen Tong
 * @Date: 2022-04-17 14:47:02
 * @LastEditTime: 2022-04-17 23:24:32
 * @Description:
 */
package main

import "fmt"

func main() {
	var i int8 = 2
	fmt.Printf("the complement of %b is: %b\n", i, ^i) // the complement of 10 is: -11

	var j uint8 = 2
	fmt.Printf("the complement of %b is: %b\n", j, ^j) // the complement of 10 is: 11111101

	fmt.Printf("the complement of %b is: %d\n", 7, 7^2)
}

// m^x
// 有符号类型m=-1, m^x:
// int8 i=2
// x = 2:  0000 0010
// m = -1: 1000 0001, m原码->补码（反码+1）: 1111 1111
// m^x: 1111 1111 ^ 0000 0010 = 1111 1101, m补码->原码（反码+1）: 1111 1101 -> 1000 0011

// 无符号类型m全部位设置为1, m^x
// uint8 i=2
// x = 2: 0000 0010
// m:     1111 1111
// m ^ x: 1111 1111 ^ 0000 0010 = 1111 1101
