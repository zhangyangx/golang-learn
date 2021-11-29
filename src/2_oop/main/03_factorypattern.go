package main

import (
	"fmt"
	"golang-study/src/2_oop/model"
)

// 工厂设计模式
func main() {
	stu := model.NewInstance(1, "zs", 22)
	fmt.Println(*stu)
	fmt.Println(stu.Id)
	fmt.Println(stu.Name)
	fmt.Println(stu.GetAge())
}
