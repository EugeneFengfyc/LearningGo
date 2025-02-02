package main

import (
	"fmt"
	"strings"
	"sync"
)

// 数据源 Goroutine
func sourceGopher(downstream chan string, wg *sync.WaitGroup) {
	defer wg.Done() // 标记任务完成
	for _, v := range []string{"Hello World", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream) // 关闭通道，通知 filerGopher 数据发送完毕
}

// 过滤器 Goroutine
func filerGopher(upstream, downstream chan string, wg *sync.WaitGroup) {
	// Done() 必须和 Add() 匹配。 每调用一次 Add(1)，必须有对应的 Done()，否则会导致死锁。
	defer wg.Done() // 标记任务完成
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream) // 关闭通道，通知 printGopher 数据处理完毕
}

// 打印器 Goroutine
func printGopher(upstream chan string, wg *sync.WaitGroup) {
	defer wg.Done() // 标记任务完成
	for item := range upstream {
		fmt.Println(item)
	}
}

func main() {

	/*sync.WaitGroup 的核心方法：
	wg.Add(n)： 增加 n 个待完成的任务。
	wg.Done()： 在任务完成时调用，减少计数器。
	wg.Wait()： 阻塞主线程，直到计数器归零，确保所有 Goroutine 执行完毕*/

	/*	main() 中：

		wg.Add(3)：表示我们有 3 个 Goroutine 需要等待。
		启动 Goroutine 后，wg.Wait() 阻塞主线程，直到所有任务完成。
		在每个 Goroutine 中：

		defer wg.Done()：确保无论函数如何退出，都会调用 Done()，减少计数器。*/

	var wg sync.WaitGroup // 初始化 WaitGroup

	c1 := make(chan string)
	c2 := make(chan string)

	// 为每个 Goroutine 增加计数器
	// Add() 必须在启动 Goroutine 之前调用。 否则可能导致 Goroutine 还未启动，Wait() 就已经结束。
	wg.Add(3)

	// 启动 Goroutine
	go sourceGopher(c1, &wg)
	go filerGopher(c1, c2, &wg)
	go printGopher(c2, &wg)

	// 等待所有 Goroutine 完成
	wg.Wait()
}
