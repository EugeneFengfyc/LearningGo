package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	} //分线路
	time.Sleep(4 * time.Second) // 主线路， 如果不在main函数里加上本行，main函数会立即返回不管sleepGopher是否执行完
}

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Snoring...", id)
}
