/*
 * @Author: Galen Tong
 * @Date: 2022-04-24 11:50:51
 * @LastEditTime: 2022-04-24 12:07:52
 * @Description: 反射
 */
package main

import (
	"fmt"
	"reflect"
)

type FLOAT float64

func main() {
	var x FLOAT = 3.4
	fmt.Println("type: ", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value: ", v)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind: ", v.Kind())
	fmt.Println("value: ", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(FLOAT)
	fmt.Println(y)
}
