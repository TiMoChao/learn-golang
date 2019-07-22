// 第一版，先简单的写一个http服务
// 服务内容主要是通过url的，读取相应的文件，输出在屏幕上
package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		// path /list/timorchao.txt  "len("/list/"):" 相当于截取/list/之后的字符串
		path := request.URL.Path[len("/list/"):]
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		// 与file.Open 对应,成对出现
		defer file.Close()

		// 去file对对应的内容
		content, err := ioutil.ReadAll(file)

		if err != nil {
			panic(err)
		}

		// 将读出来的内容写进http.Response
		writer.Write(content)
	})

	err := http.ListenAndServe(":44444", nil)
	if err != nil {
		panic(err)
	}
}
