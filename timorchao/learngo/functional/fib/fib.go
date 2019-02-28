//闭包的应用
//斐波拉契数列
//TODO 还是很不理解  为函数实现接口
//使用函数来遍历二叉树
package main

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

//TODO 还是很不离家为函数实现接口
// type fib func() int

// func fibonacci() fib {
// 	a, b := 0, 1
// 	return func() int {
// 		a, b = b, a+b
// 		return a
// 	}
// }

// //实现Read接口
// func (f fib) Read(p []byte) (n int, err error) {
// 	next := f()
// 	if next > 10000 {
// 		return 0, io.EOF
// 	}
// 	s := fmt.Sprintf("%d\n", next)

// 	//TODO: incorrect if p is too small
// 	return strings.NewReader(s).Read(p)
// }

// func printFileContents(reader io.Reader) {
// 	scanner := bufio.NewScanner(reader)
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// }

// func main() {
// 	f := fibonacci()
// 	printFileContents(f)
// }

//使用函数来遍历二叉树

func main() {

}
