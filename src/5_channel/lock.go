package main

import (
	"fmt"
	"sync"
	"time"
)

var factorialMap = make(map[int]int)

// 声明一个全局的互斥锁
var lock sync.Mutex

// 使用协程计算1~200所有数字的阶乘，并把结果放到map中
// 如果不使用lock进行加锁，就会报并发错误，因为多个协程同时操作一个变量
func main() {
	for i := 1; i <= 200; i++ {
		go factorial(i)
	}
	// 无法知道前面的线程要执行多久，随便设置10s
	time.Sleep(time.Second * 10)
	fmt.Println(factorialMap)
}

// 获取阶乘
func factorial(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 加锁
	lock.Lock()
	factorialMap[n] = res
	// 解锁
	lock.Unlock()
}
