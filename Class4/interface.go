package main

import (
	"fmt"
	"strings"
)

// 接口名称通常以er结尾
type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}

// 将laser嵌套在starship里
type starship struct {
	laser
}

// 传入接口类型
// 任何实现talker接口的类型都可以传入
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	shout(martian{})
	shout(laser(3))

	s := starship{laser(6)}
	fmt.Println(s.talk()) // 方法forwarding
	shout(s)              // 因为嵌套的laser实现了talk，starship也可以被传入
}
