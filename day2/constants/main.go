// 常量定义

package main

import "fmt"

const (
	red = iota << 1
	blue
	green
)

func main() {
	fmt.Println(red)
	fmt.Println(blue)
	fmt.Println(green)
}
