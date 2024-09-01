package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Println(i, v)
	}

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("---------------------------")
	q := [...]int{1, 2, 3}
	fmt.Printf("%T", q)

	fmt.Println("---------------------------")
	a1 := [2]int{1, 2}
	b1 := [...]int{1, 2}
	c1 := [2]int{1, 3}
	fmt.Println(a1 == b1, a1 == c1, b1 == c1) // "true false false"

	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化数组中索引为 1 和 3 的元素
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化数组中指定的元素
	array = [4][2]int{1: {0: 20, 1: 22}, 3: {1: 41}}
	fmt.Println(array)

}
