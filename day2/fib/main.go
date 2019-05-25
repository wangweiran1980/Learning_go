// 100以内的斐波那契数之和

package main

import "fmt"

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// a, b, sum := 1, 1, 0
	// for {
	// 	sum += a
	// 	a, b = b, a+b
	// 	if a >= 100 {
	// 		break
	// 	}
	// }
	// fmt.Println(sum)
	n := 1
	sum := 0
	for {
		if fib(n) >= 100 {
			break
		}
		sum += fib(n)
		n++
	}
	fmt.Println(sum)
}
