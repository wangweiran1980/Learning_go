// 显示简单html内容

package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>Hello, World!</h1>")
}

func main() {
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
}
