// 修改字符串操作, 在golang中字符串是只读的

package main

import "fmt"

func main() {
	str := "hello"
	// 第一种方式
	// 将字符串转换为byte切片
	s1 := []byte(str)
	s1[0] = 'a'
	str = string(s1)
	fmt.Println(str)

	// 第二种方式
	str = "h" + str[1:]
	fmt.Println(str)

	// 修改中文字符串
	s := "你好, 世界"
	s2 := []rune(s)
	s2[0] = 'a'
	s = string(s2)
	fmt.Println(s)
}
