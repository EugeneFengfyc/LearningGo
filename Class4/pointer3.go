package main

import "fmt"

// Go语言的函数和方法都是通过值传递参数的，因此这意味着函数总是操作于被传递参数的副本
// 当指针被传递到函数时， 函数将接受传入的内存地址的副本，之后函数通过解引用来修改指针指向的值

type stats struct {
	level  int
	health int
}

func levelUp(s *stats) {
	s.level++
}

type human struct {
	name string
	stats
}

func main() {

	player := human{
		name: "Jack",
	}

	// &不仅可以获得结构体的内存地址，还可以获取结构体中字段的地址
	levelUp(&player.stats)
	fmt.Printf("player = %+v\n", player)
	levelUp(&(player.stats))
	fmt.Printf("player = %+v\n", player)

}
