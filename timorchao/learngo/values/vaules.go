package main

import (
	"fmt"
)

func main() {

	//Strings
	fmt.Println("timorchao")
	fmt.Println("timorchao" + "你好啊")

	fmt.Println("timorchao", "你好啊")

	//Integers or Floats
	fmt.Println("3 + 5 =", 3+5)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	//Booleans
	fmt.Println(true)

	fmt.Println(false)
	fmt.Println(false && true)
	fmt.Println(false || true)
	fmt.Println(true || false)

	fmt.Println(!true)

}
