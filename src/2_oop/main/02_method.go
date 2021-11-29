package main

import "fmt"

// 方法：注意区别方法与函数
func main() {
	dog := Dog{name: "😍"}
	res := dog.getName(2)
	fmt.Println("call res:", res)

	i := integer(1)
	i.change()
	// 2
	fmt.Println(i)
}

type Dog struct {
	name string
}

// 将方法与Dog结构体绑定，方法参数age int，返回一个字符串
func (dog Dog) getName(age int) string {
	fmt.Println("This dog name is", dog.name)
	fmt.Println("This dog age is", age)
	return "success"
}

type integer int

// 此方法可以改变int的值
func (i *integer) change() {
	*i = *i + 1
}
