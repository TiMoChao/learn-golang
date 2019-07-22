package main

import (
	"errors"
	"fmt"
)

func tryRecover() {

	// 先报一个panic 如果在defer之前抛出panic，后面的defer是没有机会执行的
	//panic(errors.New("this is panic"))

	// recover 仅仅在defer中调用
	defer func() {
		// recover 返回是interface，不知道具体类型
		r := recover()
		if isErr, ok := r.(error); ok {
			fmt.Println("这是一个error:", isErr.Error())
		} else {
			panic(fmt.Sprintf("我不知道接下来做什么了：%v", r))
		}
	}()

	panic(errors.New("this is panic"))
}

func main() {
	tryRecover()
}
