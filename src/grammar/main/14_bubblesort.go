package main

import "fmt"

// 冒泡排序
func main() {
	arr := []int{2, 4, 1, 2, 5}
	bubbleSort(&arr)
}

func bubbleSort(arr *[]int) {

	fmt.Println("原始数组：", *arr)

	temp := 0

	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}

	fmt.Println("排序后数组：", *arr)
}
