package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println("loop")
		break
	}

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("----------------")

	for j := range 6 {
		fmt.Println(j)
	}
}
