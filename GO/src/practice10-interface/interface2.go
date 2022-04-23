/*
 * @Author: Galen Tong
 * @Date: 2022-04-23 22:29:56
 * @LastEditTime: 2022-04-23 22:45:24
 * @Description:多态性
 */
package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

// defines different things that have value
type valueable interface {
	getValue() float32
}

func showValue(asset valueable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	// Value of the asset is 30309.199219
	// Value of the asset is 665000.000000
	var o valueable = valueable(stockPosition{"GOOG", 7577.30, 4})
	showValue(o)
	o = car{"BMW", "M5", 665000}
	showValue(o)
}
