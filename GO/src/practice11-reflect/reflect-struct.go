/*
 * @Author: Galen Tong
 * @Date: 2022-04-24 12:09:52
 * @LastEditTime: 2022-04-24 12:21:58
 * @Description:
 */
package main

import (
	"fmt"
	"reflect"
)

type NotKnownType struct {
	s1, s2, s3 string
}

func (n NotKnownType) Add() string {
	return n.s1 + " add " + n.s2 + " add " + n.s3
}

func (n NotKnownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

var secret interface{} = NotKnownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	knd := value.Kind()
	// Ada - Go - Oberon
	// main.NotKnownType
	// struct
	fmt.Println(value)
	fmt.Println(typ)
	fmt.Println(knd)
	// Field 0: Ada
	// Field 1: Go
	// Field 2: Oberon
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}
	// Method 0:  [Ada add Go add Oberon]
	// Method 1:  [Ada - Go - Oberon]
	results1 := value.Method(0).Call(nil)
	fmt.Println("Method 0: ", results1)

	results2 := value.Method(1).Call(nil)
	fmt.Println("Method 1: ", results2)
}
