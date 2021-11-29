package main

import (
	"errors"
	"fmt"
)

// 错误处理
// 1.使用defer+recover处理错误
// 2.自定义异常
func main() {
	test()
	fmt.Println(verify(1))
}

func test() {

	// 使用defer+recover处理错误
	defer func() {
		// 内置函数recover()，可以捕获到异常
		err := recover()
		if err != nil {
			fmt.Println("err =", err)
		}
	}()
	num1 := 10
	num2 := 0
	fmt.Println(num1 / num2)
}

func verify(salary float64) (err error) {
	if salary >= 10000 {
		return nil
	} else {
		// 自定义异常
		return errors.New("工资不能小于10000")
	}
}
