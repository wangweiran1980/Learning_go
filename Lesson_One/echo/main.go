// 模拟echo -n命令行为

package main

import (
	"flag"
	"fmt"
)

func main() {
	f := flag.Bool("n", false, "取消换行符")
	flag.Parse()
	if *f {
		fmt.Print(flag.Arg(0))
	} else {
		fmt.Println(flag.Arg(0))
	}
}
