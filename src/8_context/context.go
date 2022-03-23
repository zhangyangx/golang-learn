package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 最顶层上下文
	baseCtx := context.Background()
	// 往这个上下文添加一个kv
	ctx := context.WithValue(baseCtx, "a", "b")
	// 在协程里获取上下文的kv
	go func(c context.Context) {
		fmt.Println(c.Value("a"))
	}(ctx)
	// 在最顶层上下文基础上设置一个超时上下文，超时时间1s（可以用来处理某些业务超时后就把线程中断）
	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel()
	// 开辟一个协程，根据上下文是否超时执行不同业务
	go func(c context.Context) {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			select {
			case <-c.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)
	time.Sleep(time.Second * 5)
}
