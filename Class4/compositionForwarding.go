package main

import "fmt"

// Using embedding to forward method

type report struct {
	//嵌入
	sol
	temperature
	location
}

type sol int

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (l location) days(l2 location) int {
	return 5
}

func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

func main() {
	bradbury := location{-5.234, 125.534}
	t := temperature{high: -1.0, low: -78.0}

	report := report{
		sol:         15,
		temperature: t,
		location:    bradbury,
	}

	// 此处forward了temperature的方法
	fmt.Println(report.average())
	fmt.Println(report.lat)
	fmt.Println(report.sol.days(1446))
	//fmt.Println(report.days(1446))
	//Ambiguous reference 'days' 同名嵌入方法会有歧义，调用时会报错
	//此时需要直接在report下实现方法，优先级高于嵌入类型的方法
}
