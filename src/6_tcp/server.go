package main

import (
	"fmt"
	"net"
)

// 服务端：监听端口8888
func process(conn net.Conn) {
	// 关闭conn
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("连接关闭失败")
		}
	}(conn)
	for {
		// 创建一个切片
		buf := make([]byte, 1024)
		read, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端断开连接")
			return
		}
		fmt.Print(string(buf[:read]))
	}
}

func main() {

	fmt.Println("服务端开始监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			fmt.Println("监听关闭失败")
		}
	}(listen)
	// 循环等待客户端来连接
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
		}
		go process(accept)
	}
}
