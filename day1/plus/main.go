// 命令行接受多个整数参数, 相加后返回结果
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	for _, v := range os.Args[1:] {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	fmt.Println(sum)
}
