//并发版 hello world
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	for i := 0; i < 500; i++ {
		go PrintWorld(i, ch)
	}

	//不断从管道中拿数据，但是 内容输出的go程已经结束。主程还在往管道中取数据，会报错
	//fatal error: all goroutines are asleep - deadlock!
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func PrintWorld(i int, ch chan string) {
	//针对以上报错，加一个for循环即可
	for {
		ch <- fmt.Sprintf("hello world! %d", i)
	}
}
