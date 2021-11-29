package main

import "fmt"

// 类型断言，判断变量是否属于某个类型
func main() {
	// 空接口
	var x interface{}
	f1 := 22.9
	// 空接口可以接收任意类型
	x = f1
	f := x.(float64)
	// f类型为 float64
	fmt.Printf("f类型为 %T \n", f)
	// 类型不匹配会报错，panic: interface conversion: interface {} is float64, not int
	//t := x.(int)
	//fmt.Printf("t类型为 %T", t)

	// 推荐用法，避免报错
	y, ok := x.(float64)
	if ok {
		fmt.Printf("%T \n", y)
		fmt.Println(ok)
	} else {
		fmt.Println("error")
	}
}
