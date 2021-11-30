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
// 4.使用ioutil读写文件
// 5.写文件
// 6.判断文件是否存在
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

	// 写文件
	// param1：文件路径；param2：文件打开模式；param3：权限控制
	writeFile, writeFileErr := os.OpenFile("tmp/write.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if writeFileErr != nil {
		fmt.Println("文件打开错误")
	}
	// 最后执行此行代码
	defer writeFile.Close()
	// 写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(writeFile)
	writer.WriteString("hello, world")
	// 前面写操作时写在内存中的，需要flash到文件里
	writer.Flush()

	// 使用ioutil写文件（会覆盖之前内容）
	err = ioutil.WriteFile("tmp/write.txt", content, 0666)

	exists, _ := PathExists("tmp/write.txt")
	fmt.Println(exists)
}

// PathExists 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		// 文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		// 文件不存在
		return false, nil
	}
	// 未知错误
	return false, err
}
