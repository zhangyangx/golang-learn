package main

import "fmt"

// channel：可理解为队列，先进先出
// 1.channel定义：channel之能放指定的数据类型
// 2.channel常用操作
func main() {

	// 定义一个可以存放3个int的channel
	chanInt := make(chan int, 3)
	fmt.Printf("chanInt的值：%v，chanInt地址：%p\n", chanInt, &chanInt)

	// 向channel写入数据
	chanInt <- 10
	num := 20
	chanInt <- num
	chanInt <- 22

	// 查看channel的长度和cap（容量）
	fmt.Printf("chanInt的长度：%v，chanInt的cap：%v\n", len(chanInt), cap(chanInt))

	// 遍历前需要关闭channel，关闭后不能再写入数据
	close(chanInt)
	// 遍历channel中的数据
	for v := range chanInt {
		fmt.Println(v)
	}

	//
	//for i := 0; i < cap(chanInt); i++ {
	//	fmt.Println(<-chanInt)
	//}
}
