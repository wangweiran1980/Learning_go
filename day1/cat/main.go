// 模拟cat命令, 在屏幕上打印文件

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printFile(filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}

func main() {
	for _, v := range os.Args[1:] {
		printFile(v)
	}
}
