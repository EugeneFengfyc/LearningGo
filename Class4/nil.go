package main

import (
	"fmt"
	"sort"
)

type person struct {
	age int
}

func (p *person) ageLevelUp() {
	// 避免nil导致panic
	if p == nil {
		return
	}
	p.age++
}

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool {
			return i < j
		}
	}
	sort.Slice(s, less)
}

func main() {
	var nowhere *int
	fmt.Println(nowhere)
	fmt.Println(&nowhere)

	person := person{
		age: 18,
	}
	person.ageLevelUp()
	fmt.Println(person)

	food := []string{"onion"}
	sortStrings(food, nil)
	fmt.Println(food)
}
