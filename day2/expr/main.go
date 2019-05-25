// 模拟expr程序, 只能计算两个值

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	str1 := os.Args[1]
	op := os.Args[2]
	str2 := os.Args[3]
	a, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println(err)
		return
	}

	// if op == "+" {
	// 	fmt.Println(a + b)
	// } else if op == "-" {
	// 	fmt.Println(a - b)
	// } else if op == "*" {
	// 	fmt.Println(a * b)
	// } else if op == "/" {
	// 	fmt.Println(a / b)
	// } else {
	// 	log.Fatal("操作符错误!")
	// }

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	default:
		log.Fatal("操作符错误!")
	}
}
