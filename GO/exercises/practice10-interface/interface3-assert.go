/*
 * @Author: Galen Tong
 * @Date: 2022-04-23 22:50:13
 * @LastEditTime: 2022-04-23 23:04:39
 * @Description:类型断言，检测和转换接口变量的类型
 */
package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	// The type of areaIntf is: *main.Square
	// areaIntf does not contain a variable of type Circle
	// The type *main.Square with value &{5}
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1

	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("The type %T with value %v\n", t, t)
	case Circle:
		fmt.Printf("The type %T with value %v\n", t, t)
	default:
		fmt.Println("Unexpected type %T", t)
	}
}
