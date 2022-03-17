package main

import (
	"fmt"
	"time"
)

// 单向通道
func main() {
	c := make(chan int)
	go prod(c)
	go consume(c)
	time.Sleep(time.Second * 10)
}

// 接收单向通道，只负责生产
func prod(c chan<- int) {
	for {
		c <- 1
		fmt.Println("生产...")
	}
}

// 接收单向通道，只负责消费
func consume(c <-chan int) {
	for {
		<-c
		fmt.Println("消费...")
	}
}
