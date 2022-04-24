/*
 * @Author: Galen Tong
 * @Date: 2022-04-17 14:47:02
 * @LastEditTime: 2022-04-24 13:45:14
 * @Description:
 */
package main

import "fmt"

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	duck.Quack()
	duck.Walk()
}

type Bird struct {
}

func (b *Bird) Quack() {
	fmt.Println("I am quacking!")
}

func (b *Bird) Walk() {
	fmt.Println("I am walking")
}

func main() {
	b := new(Bird)
    DuckDance(b)
}
