/*
 * @Author: Galen Tong
 * @Date: 2022-04-22 23:22:45
 * @LastEditTime: 2022-04-22 23:28:20
 * @Description: 带标签的结构体
 */
package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	tt := TagType{true, "Galen", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	println(ttType)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
