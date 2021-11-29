package main

import "fmt"

// 接口 interface
func main() {
	computer := Computer{}
	iphone := IPhone{}
	computer.Working(iphone)
	android := Android{}
	computer.Working(android)

	var usb USB = IPhone{}
	usb.Transport()
}

type USB interface {
	Transport()
	Charge()
}

type IPhone struct {
}

type Android struct {
}

type Computer struct {
}

// Transport 让IPhone实现USB方法
func (iphone IPhone) Transport() {
	fmt.Println("IPhone传输文件")
}
func (iphone IPhone) Charge() {
	fmt.Println("IPhone充电")
}

// Transport 让Android实现USB方法
func (android Android) Transport() {
	fmt.Println("Android传输文件")
}
func (android Android) Charge() {
	fmt.Println("Android充电")
}

func (computer Computer) Working(usb USB) {
	usb.Transport()
	usb.Charge()
}
