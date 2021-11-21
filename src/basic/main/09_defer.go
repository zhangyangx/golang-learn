package main

import "fmt"

// defer，将当前代码行压入栈中，最后执行（先进后出）
func main() {

	fmt.Println(sum0(1, 2))

}

func sum0(n1 int, n2 int) int {

	defer fmt.Println("param n1:", n1)
	defer fmt.Println("param n2:", n2)

	fmt.Println("hello world")

	return n1 + n2
}
