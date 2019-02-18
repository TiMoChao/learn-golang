package main

import "fmt"

func main() {
	// if i := true; i {
	if i := false; i {
		fmt.Println("我是真的")
	} else {
		fmt.Println("我是假的，哈哈哈")
	}

	if 4/2 == 2 {
		fmt.Println("4/2的结果是2哦")
	}
}
