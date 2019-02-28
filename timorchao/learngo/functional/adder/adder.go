package main

import "fmt"

//定义一个函数，返回值是一个函数，该函数为 func(int) int
func adder(sum int) func(int) int {
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder(0)
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}

}
