// error概念示例 在defer示例基础上做演示
// 循环一些数字，创建一个test.txt文件，将数字写入即可
// 写入的时候使用buffio
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "timorchao.txt"
	// 创建并且打开文件 当文件不存在的时候err != nil
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_EXCL, 0666)

	//err = errors.New("我是一个新的error")
	if err != nil {
		// 如果知道err是具体的某种错误类型，那么做具体的err处理，如果不知道，直接用err即可
		// 一般情况看注释就可以知道err属于那种err
		if pathError, ok := err.(*os.PathError); !ok {
			fmt.Printf("Errors is %s", err)
			return
		} else {
			fmt.Printf("pathError op is %s\n", pathError.Op)
			fmt.Printf("Errors path is %s\n", pathError.Path)
			fmt.Printf("Errors error is %s\n", pathError.Error())
			return
		}

		// 当有错误的时候不要panic，很不美观，需要对err进行处理
		//panic("创建打开文件" + fileName + "失败")
	}

	// 创建并打开记得关闭file，成对出现
	defer file.Close()

	// 使用bufio
	write := bufio.NewWriter(file)
	// bufio中的数据刷新到文件file中
	defer write.Flush()

	for i := 0; i < 10000; i++ {
		fmt.Fprintln(write, i)
	}
}
