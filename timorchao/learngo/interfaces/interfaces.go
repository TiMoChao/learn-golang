// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Abser interface {
// 	Abs() float64
// }

// type MyFloat float64

// type Vertex struct {
// 	X, Y float64
// }

// func main() {
// 	var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	// v := Vertex{3, 4}

// 	a = f // a MyFloat 实现了 Abser
// 	// a = &v // a *Vertex 实现了 Abser

// 	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
// 	// 所以没有实现 Abser。
// 	// a = v

// 	fmt.Println(a.Abs())
// }

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

//隐式接口
// package main

// import (
// 	"fmt"
// 	"os"
// )

// type Reader interface {
// 	Read(b []byte) (n int, err error)
// }

// type Writer interface {
// 	Write(b []byte) (n int, err error)
// }

// type ReadWriter interface {
// 	Reader
// 	Writer
// }

// func main() {
// 	var w Writer

// 	// os.Stdout 实现了 Writer
// 	w = os.Stdout

// 	fmt.Fprintf(w, "hello, writer\n")
// }

package main

import (
	"fmt"
)

type User struct {
	a int
	b string
}

type Iner string

func main() {
	a := User{123, "timorchao"}
	z := User{9001, "Zaphod Beeblebrox"}
	s := Iner("timorchao11")

	fmt.Println(a, z, s)

}

func (s User) String() string {
	return fmt.Sprintf("%v (%v years)", s.a, s.b)
}

func (s Iner) String() string {
	return fmt.Sprintf("%s", "timorchaoaaaa")
}
