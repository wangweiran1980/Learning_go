// 简单实现的md5sum功能

package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var md5Value [16]byte

func main() {
	if len(os.Args) != 2 {
		log.Fatal("参数错误!")
	}
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	md5Value := md5.Sum([]byte(data))
	fmt.Printf("%x\n", md5Value)
}
