//闭包的应用
//斐波拉契数列
//为函数实现接口
package main

import (
	"fmt"
)

// //斐波拉契数列的生成器
// //1 1 2 3 5 8 13 21 34
// //  a b
// //    a b(a 等于上一次的 b, b等于上一次的 a+b,相当于往后挪了一位)
// func fibonacci() func() int {
// 	a, b := 0, 1 //自由变量
// 	return func() int {
// 		a, b = b, a+b
// 		return a
// 	}
// }

// func main() {
// 	f := fibonacci()
// 	fmt.Println(f()) //1
// 	fmt.Println(f()) //2
// 	fmt.Println(f()) //3
// 	fmt.Println(f()) //4
// 	fmt.Println(f()) //5
// 	fmt.Println(f()) //6
// 	fmt.Println(f()) //7
// 	fmt.Println(f()) //8
// }

type adder func() int

func fibonacci() adder {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// func ()  {

// }
func main() {
	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
