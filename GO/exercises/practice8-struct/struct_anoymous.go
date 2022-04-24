/*
 * @Author: Galen Tong
 * @Date: 2022-04-22 23:34:55
 * @LastEditTime: 2022-04-22 23:41:25
 * @Description:
 */
package main

import "fmt"

type innerS struct {
	inner1 int
	inner2 int
}

type outerS struct {
	b int
	c float32
	//anonymous field
	int
	// another struct
	innerS
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 6.5
	outer.int = 60
	outer.inner1 = 5
	outer.inner2 = 10
	// &{6 6.5 60 {5 10}}
	fmt.Println(outer)
	// Literal value
	outer2 := outerS{6, 6.5, 60, innerS{5, 10}}
	// {6 6.5 60 {5 10}}
	fmt.Println(outer2)
}
