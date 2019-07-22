// defer示例
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
	// 创建并且打开文件
	file, err := os.Create(fileName)
	if err != nil {
		panic("创建打开文件" + fileName + "失败")
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
