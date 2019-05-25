// 反转字符串

package main

import "fmt"

func reverse(s string) string {
	t := []rune(s)
	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}

	return string(t)
}

func main() {
	str := "你好， 世界"
	fmt.Println(reverse(str))
}
