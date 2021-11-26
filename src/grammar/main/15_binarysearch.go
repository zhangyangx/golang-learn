package main

import (
	"fmt"
)

// 二分查找
// 前置条件：数组必须是有序的
func main() {
	ArrayNum := [6]int{1, 8, 10, 89, 1000, 1234}
	ArrayNum2 := [6]int{8, 200, 300, 889, 1000, 1234}
	BinarySearch(&ArrayNum, 0, len(ArrayNum), 1234)
	BinarySearch(&ArrayNum2, 0, len(ArrayNum), 300)
	BinarySearch(&ArrayNum, 0, len(ArrayNum), -8)
	BinarySearch(&ArrayNum2, 0, len(ArrayNum), 30)
}

func BinarySearch(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Printf("%v中找不到\t元素%v\n", *arr, findVal)
		return
	}

	// 先找到中间下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		// 说明要查找的数在左边  就应该向 leftIndex ---- (middle - 1)再次查找
		BinarySearch(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		// 如果 arr[middle] < findVal , 就应该向 middle+1---- rightIndex
		BinarySearch(arr, middle+1, rightIndex, findVal)
	} else {
		// 找到了
		fmt.Printf("%v中找到元素%v,下标为%v\n", *arr, findVal, middle)
	}

}
