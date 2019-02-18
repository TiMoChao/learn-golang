//常量
package main

import "fmt"

const a = "timorchao"
const b = 1444.44
const c = 99.9
const e = b / c

const (
	f = iota
	g
	h
	k
)

func main() {
	fmt.Println(a, b, c, e)
	fmt.Println(f, g, h, k)
}
