// 以下标为界, 反转切片的两部分

package main

import (
	"fmt"
	"log"
)

func reverseSlice(s []int, n int) []int {
	if n > len(s)-1 {
		log.Fatal("slice bounds out of range")
	}
	return append(s[n:], s[0:n]...)
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(reverseSlice(s, 3))
	fmt.Println(reverseSlice(s, 9))
}
