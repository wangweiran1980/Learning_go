// 模拟wc命令, 统计文件行数

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	paraL := flag.Bool("l", false, "统计文件行数")
	flag.Parse()
	if *paraL {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		content, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
		}
		t := string(content)
		fmt.Println(strings.Count(t, "\n"))
	}
}
