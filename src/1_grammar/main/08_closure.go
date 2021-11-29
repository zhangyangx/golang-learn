package main

import (
	"fmt"
	"strings"
)

// 闭包
func main() {
	f := addUpper()
	// 11
	fmt.Println(f(1))
	// 11+2=13,会在之前结果上累加
	fmt.Println(f(2))
	// 16
	fmt.Println(f(3))

	fileFunc := makeSuffix(".jpg")
	// 返回 大大大.jpg
	fmt.Println(fileFunc("大大大"))
}

// 累加函数,返回一个函数,n只初始化一次，每次操作都是对n的累加
func addUpper() func(int) int {
	n := 10
	// 只会打印一次
	fmt.Println("hello")
	return func(x int) int {
		n += x
		return n
	}
}

// 如果fileName以suffix结尾，则返回fileName，否则返回fileName+suffix
func makeSuffix(suffix string) func(fileName string) string {
	return func(fileName string) string {
		// strings.HasSuffix(fileName, suffix) 判断fileName是否以suffix结尾
		if strings.HasSuffix(fileName, suffix) {
			return fileName
		} else {
			return fileName + suffix
		}
	}

}
