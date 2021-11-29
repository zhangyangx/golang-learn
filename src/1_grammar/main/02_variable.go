package main

import (
	"fmt"
	"unsafe"
)

// 变量
// 1.golang中定义变量有三种方式
// 2.多变量声明
// 3.查看变量类型
// 4.字符串的两种表现形式
// 5.全局变量
var globalVariable = 1

func main() {

	// 输出全局变量
	fmt.Println(globalVariable)

	// 第一种：指定变量类型，声明后若不赋值，使用默认值
	var i int
	fmt.Println(i)

	// 第二种：根据值自行判断变量类型（类型推导）
	var num = 10.11
	fmt.Println(num)

	// 第三种：省略var
	name := "zhang yang"
	fmt.Println(name)

	// 多变量声明
	age, desc, high := 17, "hello", 1.55
	fmt.Println(age, desc, high)

	// 查看变量类型
	fmt.Printf("变量age类型%T，占用字节数为%d \n", age, unsafe.Sizeof(age))

	// 字符串表现形式一：双引号
	phoneNumber := "13408554983"
	// 输出结果：13408554983
	fmt.Println(phoneNumber)

	// 字符串表现形式二：反引号，将字符串内容原样输出
	sql := `"select * from member limit 1"`
	// 输出结果："select * from member limit 1"
	fmt.Println(sql)
}
