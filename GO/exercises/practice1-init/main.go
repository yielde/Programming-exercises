/*
 * @Author: Galen Tong
 * @Date: 2022-04-16 23:27:04
 * @LastEditTime: 2022-04-17 00:24:16
 * @Description: 初始化： pkg->const...->var...->init()->main()->Exit
 */
package main

import (
	"fmt"
	"math"
)

var Pi float64

func init() {
	Pi = 4 * math.Atan(1)
}

var twoPi = 2 * Pi

func main() {
	fmt.Println(Pi)
	fmt.Println(twoPi)
}

// result:
// 3.141592653589793
// 0
