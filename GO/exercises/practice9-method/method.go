/*
 * @Author: Galen Tong
 * @Date: 2022-04-23 00:04:18
 * @LastEditTime: 2022-04-23 00:13:33
 * @Description:结构体方法
 */
package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func main() {
	// 	sum is: 2
	// 	sum with param is: 5
	// 	sum is: 3
	// 	sum with param is: 6
	two1 := new(TwoInts)
	two1.a = 1
	two1.b = 1
	fmt.Printf("sum is: %d \n", two1.AddThem())
	fmt.Printf("sum with param is: %d \n", two1.AddToParam(3))
	two2 := TwoInts{1, 2}
	fmt.Printf("sum is: %d \n", two2.AddThem())
	fmt.Printf("sum with param is: %d \n", two2.AddToParam(3))
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
