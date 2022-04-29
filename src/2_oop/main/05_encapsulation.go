package main

import "golang-learn/src/2_oop/model"

// 封装
func main() {
	stu := model.NewInstance(1, "rose", 17)
	stu.GetAge()
}
