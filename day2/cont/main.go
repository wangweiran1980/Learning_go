// 接收标准输入, 统计字符数, 字数, 行数

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	t := []rune(string(buf))
	fmt.Println(len(string(buf)))
	fmt.Println(len(t))
	fmt.Println(strings.Count(string(buf), "\n"))
}
