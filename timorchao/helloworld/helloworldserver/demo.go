//hello world 服务器版本
package main

import (
	"fmt"
	"net/http"
)

type helloHandle struct{}

func main() {
	http.Handle("/", &helloHandle{})
	http.ListenAndServe(":12345", nil)
}

//实现了Handler的ServeHTTP方法
func (h *helloHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1><center>你好,%s,我是%s孩</center></h1>", r.FormValue("name"), r.FormValue("sex"))
	// w.Write([]byte("hello world"))
}
