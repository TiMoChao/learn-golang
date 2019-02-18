package main

import (
	"fmt"
)

func main() {
	a := make(map[int]string)

	a[1] = "hello world"
	a[2] = "timorchao"
	a[3] = "你好吗"

	//use range of for
	for k, v := range a {
		fmt.Println("k:", k, "v:", v)
	}

	// b, ok := a[99]
	b, ok := a[1]
	fmt.Println(b, ok)

	b, ok = a[3]
	fmt.Println(b, ok)

	delete(a, 3)

	b, ok = a[3]
	fmt.Println(b, ok)

	//while
	for l := 0; l < 10; {
		fmt.Println(l)
		l++
	}

	k := 11
	for k < 20 {
		fmt.Println(k)
		k++
	}

	//死循环，很多时候单配channel使用
	// u := 0
	// for {
	// 	fmt.Println(u)
	// 	u++
	// }

	for n := 0; n < 10000; n++ {
		fmt.Println(n)
	}
}
