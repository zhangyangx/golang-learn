package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件常用操作
// 1.打开文件
// 2.读取文件
// 3.关闭文件
func main() {

	// 打开文件
	file, err := os.Open("tmp/hello.txt")
	if err != nil {
		fmt.Println("err:", err)
	}
	// 输出文件 输出结果是一个指针{0xc0000a4780}
	fmt.Println(*file)

	// 读取文件
	// 创建缓冲区
	const (
		defaultBufSize = 4096
	)
	reader := bufio.NewReader(file)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		// EOF表示文件的末尾，表示文件读取完毕
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}
	fmt.Println("hello.txt文件读取完毕")

	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("err:", err)
	}

	// 使用ioutil一次性将整个文件读取到内存中(适用于文件不大的情况下)
	content, e := ioutil.ReadFile("tmp/songlist.txt")
	if e != nil {
		fmt.Println("文件读取失败")
	}
	fmt.Println(string(content))

}
