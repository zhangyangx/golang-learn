package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	// 开启一个协程
	go func() {
		fmt.Println("hello from goroutine")
		// 数据写入channel
		ch <- 0
	}()
	// 主线程从channel中读取数据
	i := <-ch
	fmt.Println("i:", i)
}
