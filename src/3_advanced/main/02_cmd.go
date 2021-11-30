package main

import (
	"flag"
	"fmt"
)

// 命令行参数解析
func main() {

	// 定义几个变量，用于接收命令行参数
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认localhost")
	flag.IntVar(&port, "p", 3306, "端口，默认3306")
	flag.Parse()
	// 输出结果
	fmt.Printf("user:%v，pwd:%v,host:%v,port:%v", user, pwd, host, port)
}
