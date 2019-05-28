// 数组和切片

package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [...]int{1, 2, 3}
	b := []int{4, 5, 6, 7, 8}
	copy(a[:], b)
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	c := [...]int{4, 5, 6}
	if reflect.DeepEqual(a, c) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	d := make([]int, 0, 5)
	fmt.Printf("len = %d cap = %d\n", len(d), cap(d))
}
