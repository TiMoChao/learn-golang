// 第四班版,加上panic和recover， 区分error，有些是可以给用户看见的，有些是不可以的
// 服务内容主要是通过url的，读取相应的文件，输出在屏幕上
// error封装采用了函数式编程，函数是一等公民，可以作为参数传递
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 错误处理
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 处理panic
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic:%v\n", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
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

// 定义接口
type userError interface {
	error
	Message() string
}

func main() {
	// 故意将、/list/ 改成 / 让其能够抛出panic
	//http.HandleFunc(prefix, errWrapper(HandlerFileListfunc))
	http.HandleFunc("/", errWrapper(HandlerFileListfunc))

	err := http.ListenAndServe(":44444", nil)
	if err != nil {
		panic(err)
	}
}

const prefix = "/list/"

// 用的人自己管自己的userError
type userEr string

func (e userEr) Error() string {
	return e.Message()
}

func (e userEr) Message() string {
	return string(e)
}

// 将业务逻辑抽离成一个方法
func HandlerFileListfunc(writer http.ResponseWriter, request *http.Request) error {
	// 校验url是否以/list/开头
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userEr("path must start with " + prefix)
	}

	// path /list/timorchao.txt  "len("/list/"):" 相当于截取/list/之后的字符串
	path := request.URL.Path[len(prefix):]
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
