package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// 协程：轻量级线程
// 使用defer+recover，解决万一协程中发生异常，导致程序崩溃的问题
func main() {

	// 设置cpu数(go1.8后不用设置了)
	// 获取当前系统cpu数量
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println("当前系统cpu数量：", num)

	go test()
	// 避免主线程结束
	time.Sleep(time.Second * 100)
}

// 每隔一秒输出hello world
func test() {

	// 使用defer+recover，解决万一协程中发生异常，导致程序崩溃的问题
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发生异常")
		}
	}()
	var myMap map[int]int
	myMap[0] = 1
	for i := 0; i < 10; i++ {
		fmt.Println("hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
