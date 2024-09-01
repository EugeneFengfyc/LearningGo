package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	FlagNone = 1 << iota // iota 使用了一个移位操作，每次将上一次的值左移一位（二进制位），以得出每一位的常量值。
	FlagRed
	FlagGreen
	FlagBlue
)

// 声明芯片类型
type ChipType int

const (
	None ChipType = iota
	CPU           // 中央处理器
	GPU           // 图形处理器
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}

	return "N/A"
}

func main() {
	var today Weekday = Sunday
	fmt.Println(today)

	fmt.Println("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
	fmt.Println("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)

	// 输出CPU的值并以整型格式显示
	// String() 方法的 ChipType 在使用上和普通的常量没有区别。当这个类型需要显示为字符串时，Go语言会自动寻找 String() 方法并进行调用。
	fmt.Println("%s %d", CPU, CPU)
}
