package main

import (
	"fmt"
	"sync"
)

// Mutual Exclusive

type Visited struct {
	mutex   sync.Mutex
	Visited map[string]int
}

func (v *Visited) VisitLink(url string) int {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	count := v.Visited[url]
	count++
	v.Visited[url] = count
	return count
}

func main() {
	/*	误区：mutex 保护的是数据本身。
		事实：mutex 只是防止同一时间多个 goroutine 进入某个代码块，数据本身没有“自动防护”。

		误区：加锁和解锁的位置无所谓。
		事实：Lock 和 Unlock 的范围决定了保护的数据范围，错误的位置会导致竞态条件或死锁。

		误区：每个字段都需要一个锁。
		事实：不必要的多锁会增加复杂性，合理的锁粒度（比如保护整个临界区）更为高效。*/
	v := &Visited{Visited: make(map[string]int)}
	fmt.Println(v.VisitLink("https://example.com")) // 输出：1
	fmt.Println(v.VisitLink("https://example.com")) // 输出：2
	fmt.Println(v.VisitLink("https://golang.org"))  // 输出：1

}
