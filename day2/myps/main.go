// 模拟ps命令, 显示pid和cmdline

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	infos, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, info := range infos {
		if info.IsDir() {
			fmt.Println(info.Name())
		}
	}
}
