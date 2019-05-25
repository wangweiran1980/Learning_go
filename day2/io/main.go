// 标准输入、标准输出

package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(buf))
	os.Stdout.Write(buf)
}
