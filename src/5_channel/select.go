package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int, 5)
	c2 := make(chan int, 5)
	for i := 0; i < 10; i++ {
		i := i
		go func() { c1 <- i }()
		go func() { c2 <- i }()
	}
	select {
	case v := <-c1:
		fmt.Println("c1:", v)
	case v := <-c2:
		fmt.Println("c2:", v)
	default:
		fmt.Println("...")
	}
	time.Sleep(time.Second * 10)
}
