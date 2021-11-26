package main

import (
	"fmt"
	"strconv"
)

// 数据类型转换
func main() {

	// int32转int64
	var i1 int32 = 13
	var i2 int64 = int64(i1)

	// int64转float64
	var f1 float64 = float64(i2)

	fmt.Printf("i1=%v i2=%v fl=%v\n", i1, i2, f1)

	// 基本数据类型转string
	// 方式一：fmt.Sprintf("%参数", 表达式)
	var s1 = fmt.Sprintf("%d", i1)
	var s2 = fmt.Sprintf("%f", f1)
	fmt.Printf("s1 type %T, s1=%q \n", s1, s1)
	fmt.Printf("s2 type %T, s2=%q \n", s2, s2)

	// 方式二：使用strconv包的函数
	// 基础类型转string（param1：变量，param2：进制）
	var s3 = strconv.FormatInt(i2, 10)
	fmt.Printf("s3 type %T, s3=%v \n", s3, s3)
	// string转基础类型
	var f2, _ = strconv.ParseFloat(s2, 64)
	fmt.Printf("f2 type %T, f2=%v \n", f2, f2)

}
