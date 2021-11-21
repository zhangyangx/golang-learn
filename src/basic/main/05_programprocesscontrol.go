package main

import (
	"fmt"
	"math/rand"
)

// 程序流程控制：顺序、分支、循环
// 1.顺序控制：代码由上到下执行，没有任何判断和跳转
// 2.分支控制：让程序有选择执行
// 3.循环控制：一段代码循环执行
func main() {

	// 返回一个随机的整数 n，0 <= n <= 100。
	i := rand.Intn(100)
	fmt.Println("i =", i)

	// 分支控制
	if i < 20 {
		fmt.Println("i小于20")
	} else if i < 50 {
		fmt.Println("i大于20小于50")
	} else {
		fmt.Println("i大于50")
	}

	// switch分支控制
	switch i {
	case 20, 30, 40:
		fmt.Println("i取值为20、30或者40")
	case 50:
		fmt.Println("i取值为50")
	default:
		fmt.Println("没有匹配数据")
	}

	// type-switch：判断某个interface变量中实际指向的变量类型
	var x interface{}
	y := 10.0
	x = y
	switch z := x.(type) {
	case nil:
		fmt.Printf("x类型：%T", z)
	case int:
		fmt.Printf("x类型：%T", z)
	case float64:
		fmt.Printf("x类型：%T", z)
	case func(int):
		fmt.Printf("x类型：%T", z)
	case bool:
		fmt.Printf("x类型：%T", z)
	case string:
		fmt.Printf("x类型：%T", z)
	}
	fmt.Println()

	// 循环控制
	for i := 0; i < 3; i++ {
		fmt.Println("循环控制语句", i)
	}

	// for-range，用来遍历字符串
	str := "hello"
	for index, val := range str {
		fmt.Printf("index %d,val %c \n", index, val)
	}

}
