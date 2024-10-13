package main

import "fmt"

func main() {

	var map1 map[string]int
	map1 = map[string]int{"one": 1, "two": 2}
	fmt.Println(map1)

	var map2 map[string]int
	map2 = map1
	map2["two"] = 22
	fmt.Println(map1)

	// make(map[keytype]valuetype, cap)
	var map3 = make(map[string]int)
	map3["one"] = 1
	fmt.Println(map3)

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	delete(map1, "one")
	fmt.Println(map1)
	map1["three"] = 3

	/*
		注意：可以使用 make()，但不能使用 new() 来构造 map，
		如果错误的使用 new() 分配了一个引用对象，会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：
	*/

}
