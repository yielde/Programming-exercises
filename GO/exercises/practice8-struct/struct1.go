/*
 * @Author: Galen Tong
 * @Date: 2022-04-22 22:14:12
 * @LastEditTime: 2022-04-22 22:55:42
 * @Description:
 */
package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// The name of the person is GALEN TONG
	// The name of the person is GALEN TONG
	// The name of the person is GALEN TONG

	//1 struct as a value type
	var pers1 Person
	pers1.firstName = "Galen"
	pers1.lastName = "Tong"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s \n", pers1.firstName, pers1.lastName)

	//2 struct as a pointer
	pers2 := new(Person)
	pers2.firstName = "Galen"
	pers2.lastName = "Tong"
	(*pers2).lastName = "Tong"
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s \n", pers2.firstName, pers2.lastName)

	//3 struct as a literal
	pers3 := &Person{"Galen", "Tong"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s \n", pers3.firstName, pers3.lastName)
}
