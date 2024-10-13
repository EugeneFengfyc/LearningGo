package main

import "fmt"

func main() {
	var a1 []int
	a1 = append(a1, 1)
	fmt.Println(a1)

	a1 = append(a1, 1, 2, 3)
	fmt.Println(a1)

	a1 = append(a1, []int{1, 2, 3}...)
	fmt.Println(a1)

	/*
		在使用 append() 函数为切片动态添加元素时，如果空间不足以容纳足够多的元素，切片就会进行“扩容”，此时新切片的长度会发生改变。
		切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16……
		切片长度 len 并不等于切片的容量 cap。
	*/
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

	/*
		除了在切片的尾部追加，我们还可以在切片的开头添加元素：
		在切片开头添加元素一般都会导致内存的重新分配，而且会导致已有元素全部被复制 1 次，因此，从切片的开头添加元素的性能要比从尾部追加元素的性能差很多。
	*/
	var a2 []int
	a2 = append([]int{0}, a2...)
	fmt.Println(a2)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)
	fmt.Println(slice2)

	var a3 = []int{1, 2, 3}
	a3 = a3[1:] // 删除开头1个元素
	a3 = append(a3[:0], a3[1:]...)

	var a4 = []int{1, 2, 3, 4, 5, 6}
	a4 = append(a4[:1], a4[2:]...) // 删除中间1个元素
	fmt.Println(a4)

	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	// 迭代每一个元素，并显示其值
	for index, value := range slice {
		// 需要强调的是，range 返回的是每个元素的副本，而不是直接返回对该元素的引用
		fmt.Printf("Index: %d Value: %d Value-Addr: %X ElemAddr: %X \n", index, value, &value, &slice[index])
	}
}
