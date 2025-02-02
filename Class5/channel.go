package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 通道 c 是无缓冲的，意味着发送和接收操作都会阻塞，直到双方准备好，这确保了每个 goroutine 完成后主线程才能继续。
	c := make(chan int)
	for i := 0; i < 5; i++ {
		// 尽管每个 goroutine 都“睡”了 3 秒，整个程序的总执行时间大约也是 3 秒，而不是 15 秒（5 个 * 3 秒），因为它们是并行等待的。
		go sleepy(i, c)
	} //分线路
	for i := 0; i < 5; i++ {
		gopherId := <-c
		fmt.Println("gopherId:", gopherId, " has finished sleeping")
	}

	d := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyRandom(i, d)
	}

	timeout := time.After(2 * time.Second) // time.After返回一个channel， 在一定时间后这个channel会接收到一个值
	for i := 0; i < 5; i++ {
		select { // select语句在不包含任何case的情况下会一直等待下去
		// select的每个case都有一个通道用来发送或者接受数据， select会等待直到某个case分支的操作就绪就会执行当前的case
		case randomSleepId := <-d:
			fmt.Println("randomSleepId:", randomSleepId)
		case <-timeout:
			// 即时停止等待goroutine，只要main函数还没有返回，仍在运行的goroutine会一直占用内存
			fmt.Println("timeout")
			return
		}
	}

}

func sleepy(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Snoring...", id)
	c <- id
}

func sleepyRandom(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond) // 随机生成一个0～4秒内的时间
	c <- id
}
