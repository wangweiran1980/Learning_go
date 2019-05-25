// 九九乘法表

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("day2/multiplication_tables/tables.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fprintf(f, "%d * %d = %-4d", j, i, j*i)
		}
		fmt.Fprintln(f)
	}
}
