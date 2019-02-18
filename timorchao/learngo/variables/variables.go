package main

import "fmt"

func main() {
	// `var` declares 1 or more variables

	//显式声明类型
	var a int = 10
	var b string = "timorhcao"
	var d float32 = 123.243

	//隐式变量类型，编译器会自动根据后面的值识别
	var e = 10
	var f = "timorchao"
	var g = 123.243 //系别系统默认位数为64位。 隐式类型为: float64

	//打印
	fmt.Printf("类型为:%T, 数值为:%v\n", a, a)
	fmt.Printf("类型为:%T, 数值为:%v\n", b, b)
	fmt.Printf("类型为:%T, 数值为:%v\n", d, d)
	fmt.Printf("类型为:%T, 数值为:%v\n", e, e)
	fmt.Printf("类型为:%T, 数值为:%v\n", f, f)
	fmt.Printf("类型为:%T, 数值为:%v\n", g, g)

	var h, i, j, k = 12, "hello world", 987.32, false
	// 不声明变量类型可以同时赋值多个不同类型的变量，如果声明了变量类型，那么只能同时赋值同一种类型的数值
	// var l, m, n, o  int string float64 bool = 12, "hello world", 987.32, false
	var l, m, n, o int = 12, 13, 14, 15

	fmt.Printf("类型为:%T, 数值为:%v\n", h, h)
	fmt.Printf("类型为:%T, 数值为:%v\n", i, i)
	fmt.Printf("类型为:%T, 数值为:%v\n", j, j)
	fmt.Printf("类型为:%T, 数值为:%v\n", k, k)
	fmt.Printf("类型为:%T, 数值为:%v\n", l, l)
	fmt.Printf("类型为:%T, 数值为:%v\n", m, m)
	fmt.Printf("类型为:%T, 数值为:%v\n", n, n)
	fmt.Printf("类型为:%T, 数值为:%v\n", o, o)

	//no use `var`
	p := false
	q := 144.1444
	r, s, k := "你好么，罗超", 45, true

	fmt.Printf("类型为:%T, 数值为:%v\n", p, p)
	fmt.Printf("类型为:%T, 数值为:%v\n", q, q)
	fmt.Printf("类型为:%T, 数值为:%v\n", r, r)
	fmt.Printf("类型为:%T, 数值为:%v\n", s, s)
	fmt.Printf("类型为:%T, 数值为:%v\n", k, k)

}
