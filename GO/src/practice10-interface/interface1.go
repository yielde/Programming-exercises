/*
 * @Author: Galen Tong
 * @Date: 2022-04-23 22:13:31
 * @LastEditTime: 2022-04-23 22:35:39
 * @Description:多态，接口
 */
package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	// Looping through shapes for area ...
	// Shape details:  {5 3}
	// Area of this shape is:  15
	// Shape details:  &{5}
	// Area of this shape is:  25
	rect := Rectangle{5, 3}
	square := &Square{5}

	shapes := []Shaper{Shaper(rect), Shaper(square)}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
