package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 客户端：连接服务器的端口8888
// 客户端可以发送单行数据，然后就退出
// 在终端输入exit，表示退出
func main() {

	conn, err := net.Dial("tcp", "192.168.113.1:8888")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}
	// 发送单行数据，然后退出
	// os.Stdin代表标准输入[终端]
	reader := bufio.NewReader(os.Stdin)
	// 从终端读取一行用户输入，并准备发送给服务器
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readSting() err:", err)
	}
	// 将line发送给服务器
	write, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("conn write err:", err)
	}

	fmt.Printf("客户端发送了%d字节的数据", write)

}
