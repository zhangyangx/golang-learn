package main

import "fmt"

// 结构体
// 1.结构体定义
// 2.四种方式声明结构体
func main() {
	// 声明方式一
	var cat Cat
	cat.age = 1
	cat.name = "蛋蛋"
	// cat = {1 蛋蛋}
	fmt.Println("cat =", cat)

	var newCat Cat
	newCat = cat
	newCat.name = "露露"
	// cat = {1 蛋蛋}
	fmt.Println("cat =", cat)
	// new cat = {1 露露}
	fmt.Println("new cat =", newCat)

	var people People
	people.age = 23
	people.name = "kzlh"
	people.cat = newCat
	// kzlh: {23 kzlh {1 露露}}
	fmt.Println("kzlh:", people)

	// 声明方式二
	firstCat := Cat{age: 2, name: "🐖"}
	fmt.Println("firstCat:", firstCat)

	// 声明方式三：结构体指针
	var p1 *People = new(People)
	(*p1).age = 18
	// 底层做了处理，可以不加*
	p1.name = "zs"
	// &{18 zs {0 }}
	fmt.Println(*p1)

	// 声明方式四：结构体指针
	var p2 *People = &People{age: 17}
	p2.name = "l"
	// {17 l {0 }}
	fmt.Println(*p2)

}

// Cat 定义结构体
type Cat struct {
	age  int
	name string
}
type People struct {
	age  int
	name string
	// 结构体变量
	cat Cat
}
