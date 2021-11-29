package main

import "fmt"

// 数组
// 1.数组定义
// 2.数组遍历
func main() {

	// 定义数组
	var array [3]float64
	// 为数组复制
	array[0] = 1
	array[1] = 3
	array[2] = 5
	fmt.Println(array)
	// 数组在内存中的地址
	fmt.Printf("array的地址=%p，array[0]的地址=%p，array[1]的地址=%p，array[2]的地址=%p", &array, &array[0], &array[1], &array[2])

	// 四种定义数组的方式
	var arr1 = [3]int{2, 12, 4}
	fmt.Println(arr1)
	var arr2 = [...]int{2, 3, 4}
	fmt.Println(arr2)
	var arr3 = [...]int{1: 200, 0: 222, 2: 211}
	fmt.Println(arr3)
	arr4 := [...]string{1: "hello", 2: "world", 0: "~"}
	fmt.Println(arr4)

	// 数组遍历
	for index, value := range array {
		fmt.Println(index, "--", value)
	}

}
