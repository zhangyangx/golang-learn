package main

import "golang-study/src/2_oop/model"

// 封装
func main() {
	stu := model.NewInstance(1, "rose", 17)
	stu.GetAge()
}
