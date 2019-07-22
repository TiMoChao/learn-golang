// 第三版,开始包装error，先简单的写一个http服务
// 服务内容主要是通过url的，读取相应的文件，输出在屏幕上
// error封装采用了函数式编程，函数是一等公民，可以作为参数传递
package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 错误处理
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(HandlerFileListfunc))

	err := http.ListenAndServe(":44444", nil)
	if err != nil {
		panic(err)
	}
}

// 将业务逻辑抽离成一个方法
func HandlerFileListfunc(writer http.ResponseWriter, request *http.Request) error {
	// path /list/timorchao.txt  "len("/list/"):" 相当于截取/list/之后的字符串
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	// 与file.Open 对应,成对出现
	defer file.Close()

	// 去file对对应的内容
	content, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	// 将读出来的内容写进http.Response
	writer.Write(content)

	return nil
}
