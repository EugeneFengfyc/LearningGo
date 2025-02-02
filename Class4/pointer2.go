package main

import (
	"fmt"
	"reflect"
)

func main() {
	answer := 43
	fmt.Println("Address of answer = ", &answer)

	// Go语言不允许地址的直接修改如address++
	address := &answer
	fmt.Println("Value of address = ", *address)
	fmt.Println("Type of address = ", reflect.TypeOf(address)) //指向int类型的指针 *int
	fmt.Printf("Type of address is a %T\n", address)

	var stringPointer *string

	string1 := "string1"
	stringPointer = &string1
	fmt.Println(*stringPointer)

	string2 := "string2"
	stringPointer = &string2
	fmt.Println(*stringPointer)

	string2 = "string string"
	fmt.Println(*stringPointer)

	*stringPointer = "string string2"
	fmt.Println(string2)

	type person struct {
		name, power string
		age         int
	}

	personPointer := &person{
		name: "Jason",
		age:  8,
	}

	// 效果相同,自动解引用
	personPointer.power = "strong"
	(*personPointer).power = "strong"

	fmt.Printf("%+v\n", *personPointer)

	// 数组切片操作会自动解引用
	arrayPointer := &[3]int{1, 2, 3}
	fmt.Println(arrayPointer[0])
	fmt.Println(arrayPointer[1:2])
}
