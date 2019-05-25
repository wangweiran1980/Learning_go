// 字符串转换数字

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "123"
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}
