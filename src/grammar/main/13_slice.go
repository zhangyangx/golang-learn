package main

import "fmt"

// 切片：与数组类似，数组是值类型，切片是引用类型；切片是可以动态变化的数组
// 切片三种定义方式
func main() {

	array := [...]int{11, 32, 412, 213, 222}
	// 定义方式一：剪切数组，slice起始下标为1，最后的下表为3，但不包含3
	slice1 := array[1:3]
	// [32 412]
	fmt.Println(slice1)

	// 方式二：var 切片名 []type = make([]type,len,cap)
	var slice2 []float64 = make([]float64, 5, 10)
	slice2[1] = 10
	slice2[3] = 20
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// 方式三：直接指定具体数组
	var slice3 []string = []string{"tom", "jack", "mary"}
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
