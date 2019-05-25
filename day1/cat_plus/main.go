// 能够在屏幕上打印本地文本文件和html

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func printFile(filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}

func printHTML(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(res))
}

func main() {
	for _, name := range os.Args[1:] {
		if strings.HasPrefix(name, "http://") {
			printHTML(name)
		} else {
			printFile(name)
		}
	}
}
