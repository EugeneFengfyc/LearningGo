package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "process mode")

func swap(a, b *int) {
	/*
		*操作符作为右值时，意义是取指针的值，
		作为左值时，也就是放在赋值操作符的左边时，表示 a 指针指向的变量。其实归纳起来，*操作符的根本意义就是操作指针指向的变量。
		当操作在右值时，就是取指向变量的值，当操作在左值时，就是将值设置给指向的变量。
	*/
	t := *a
	*a = *b
	*b = t
}

func main() {
	var v1 int = 1
	var v2 string = "v2"
	fmt.Println("%p, %p", &v1, &v2)

	fmt.Println("-----------------------------------")
	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)
	fmt.Println("-----------------------------------")

	// 准备两个变量, 赋值1和2
	x, y := 1, 2
	// 交换变量值
	swap(&x, &y)
	// 输出变量值
	fmt.Println(x, y)

	flag.Parse()
	fmt.Println(*mode)

	// 使用new()创建指针
	str := new(string)
	*str = "new pointer"
	fmt.Println(*str)

	v3 := 25
	var v4 *int
	if v4 == nil {
		fmt.Println("v4 is", v4)
		v4 = &v3
		fmt.Println("v4 after initialization is", v4)
	}
}
