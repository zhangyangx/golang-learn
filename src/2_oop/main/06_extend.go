package main

import (
	"fmt"
	"golang-study/src/2_oop/model"
)

// 继承
func main() {
	adult := model.Adult{}
	adult.Name = "张三"
	adult.SetAge(20)
	adult.FindJob("coder")
	fmt.Println(adult)
}
