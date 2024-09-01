package main

import "fmt"
import "math"

func main() {

	var k1 int
	k1 = 29
	fmt.Println(k1)

	var k2 int = 30
	fmt.Println(k2)

	k3 := "k3"
	fmt.Println(k3)

	var k4, k5 int = 4, 5
	fmt.Println(k4, k5)

	var k6, k7 string = "k6", "k7"
	fmt.Println(k6, k7)

	var k8 int
	fmt.Println(k8)

	var (
		k9  = 9
		k10 = "k10"
		k11 int
	)
	fmt.Println(k9, k10, k11)

	k12, k13 := 1, "k13"
	fmt.Println(k12, k13)

	fmt.Println("------------------------")
	fmt.Println("Short hand syntax can only be used when at least one of the variables on the left side of := is newly declared. Consider the following program")
	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)
	fmt.Println("------------------------")

	k14, k15 := 14, 15
	k16 := math.Min(float64(k14), float64(k15))
	fmt.Println(k16)

}
