package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// 逻辑中使用的某个变量
	count int
	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
)

func main() {
	count = 1000
	for i := 0; i < 1000; i++ {
		go func() {
			countGuard.Lock()
			count--
			countGuard.Unlock()
		}()
	}
	time.Sleep(time.Second * 10)
	fmt.Println(count)
}
