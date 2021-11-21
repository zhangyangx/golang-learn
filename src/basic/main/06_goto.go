package main

import "fmt"

// goto：无条件地转移到程序中指定的行
func main() {

	fmt.Println(1)
	goto label
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
label:
	fmt.Println(5)
}
