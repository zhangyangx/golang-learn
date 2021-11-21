package main

import "fmt"

// 指针
// 1.获取内存地址
// 2.指针变量
func main() {

	// 获取变量的地址，使用&
	var i = 32
	// 0xc00000a098
	fmt.Println(&i)

	// 指针变量
	var ptr *int = &i
	// 0xc00000a098
	fmt.Println("ptr =", ptr)
	// 0xc000006030
	fmt.Println("ptr的地址 =", &ptr)
	// 32
	fmt.Println("ptr指向的值 =", *ptr)

	// 修改ptr指向的值
	*ptr = 10
	fmt.Println(i)

}
