package main

import (
	"fmt"
	"golang-study/src/basic/util"
)

// 函数
// func 函数名(形参列表) (返回值列表) {
//		执行语句
//		return 返回值列表
// }
// 1.函数定义
// 2.调用其他包函数
// 3.init函数
// 4.函数也是数据类型，可作为函数参数传递
// 5.匿名函数
// 6.全局函数

var (
	// GlobalFun 全局匿名函数
	GlobalFun = func(n1 int, n2 int) int {
		return n1 + n2
	}
)

func main() {
	// 调用工具包下的计算函数
	fmt.Println(util.Cal(1114, 22, '/'))

	// 函数类型变量
	a := util.Cal
	// 调用函数变量
	fmt.Println(a(3, 1, '+'))
	// a type:func(float64, float64, uint8) float64
	fmt.Printf("a type:%T\n", a)

	// 函数可做参数传递
	a2 := sum(a, 2, 3, '-')
	fmt.Println("a2:", a2)

	// 匿名函数1
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(2, 3)
	fmt.Println("res1:", res1)
	// 匿名函数2
	res2 := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Println("res2:", res2(31, 1))

	// 调用全局匿名函数
	fmt.Println("conFun:", GlobalFun(22, 22))
}

// init函数，每个源文件都可以包含一个init函数，该函数会在main函数之前执行
// 调用优先级：全局变量定义>init函数>main函数
func init() {
	fmt.Println("程序启动..........")
}

func sum(funVar func(float64, float64, byte) float64, num1 float64, num2 float64, operator byte) float64 {
	return funVar(num1, num2, operator)
}
