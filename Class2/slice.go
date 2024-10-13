package main

import "fmt"

func main() {
	/*
		切片（slice）是对数组的一个连续片段的引用，所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型），
		这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集，需要注意的是，终止索引标识的项不包括在切片内。
	*/
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])

	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])

	// 把切片的开始和结束位置都设为 0 时，生成的切片将变空
	aa := []int{1, 2, 3}
	fmt.Println(aa[0:0])

	// 如果需要动态地创建一个切片，可以使用 make() 内建函数
	// make( []Type, size, cap )
	a1 := make([]int, 2)
	b1 := make([]int, 2, 10)
	fmt.Println(a1, b1)
	fmt.Println(len(a1), len(b1))
	/*
		其中 a 和 b 均是预分配 2 个元素的切片，只是 b 的内部存储空间已经分配了 10 个，但实际使用了 2 个元素。
		容量不会影响当前的元素个数，因此 a 和 b 取 len 都是 2。
	*/

	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)
}
