package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串相关函数
func main() {

	str := "哈喽，沃德"

	// 字符串长度
	fmt.Println("str长度:", len(str))
	// 字符串遍历，可解决中文问题
	arr := []rune(str)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%c\n", arr[i])
	}
	// 字符串转整数
	n, _ := strconv.Atoi("12")
	fmt.Printf("n type:%T\n", n)
	// 整数转字符串
	s := strconv.Itoa(12345)
	fmt.Printf("s type:%T\n", s)
	// 字符串转byte[]
	bytes := []byte(str)
	fmt.Printf("bytes[] value:%v\n", bytes)
	// byte[]转字符串
	str = string([]byte{229, 147, 136, 229, 150, 189, 239, 188, 140, 230, 178, 131, 229, 190, 183})
	fmt.Printf("str value:%v\n", str)
	// 十进制转二、八、十六进制
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123 二进制:%v\n", str)
	str = strconv.FormatInt(123, 8)
	fmt.Printf("123 八进制:%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123 十六进制:%v\n", str)
	str = "哈喽，沃德"
	// 查找字符串是否包含在指定字符串中
	fmt.Println(`'哈喽'包含在str中:`, strings.Contains(str, "哈喽x"))
	// 统计字符串有几个指定的字符串
	fmt.Println(`’hello world‘中有`, strings.Count("hello world", "ll"), `个'll'`)
	// 字符串比较 EqualFold不区分大小写，==区分大小写
	fmt.Println(`'aa'等于'aA':`, strings.EqualFold("aa", "aA"))
	// 子字符串在字符串中第一次出现的index,没有则返回-1
	fmt.Println(`'da'在'adsdadaxcs'中第一次出现的index:`, strings.Index("adsdadaxcs", "da"))
	// 子字符串在字符串中最后一次出现的index,没有则返回-1
	fmt.Println(`'da'在'adsdadaxcs'中最后一次出现的index:`, strings.LastIndex("adsdadaxcs", "da"))
	// 字符串替换,param1:源字符串,param2:要替换的旧字符串,param3:要替换成的新字符串,param4:替换个数，-1表示全部替换
	// 替换全部 n=-1 或者 使用ReplaceAll
	fmt.Println(strings.Replace("Java比Go👌", "Go", "Golang", 1))
	// 字符串拆分成字符串数组
	fmt.Println(strings.Split("Hello World", " "))
	// 大小写转换
	fmt.Println(strings.ToLower("AAA"))
	fmt.Println(strings.ToUpper("AaA"))
}
