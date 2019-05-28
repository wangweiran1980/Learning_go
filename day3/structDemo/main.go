package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Student 定义学生信息
type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// stu1 := Student{ID: "0123", Name: "Jackon", Age: 20}
	// jsonData, err := json.Marshal(stu1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // 以json形式写入文件
	// f, err := os.Create("day3/structDemo/info.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// if _, err = fmt.Fprintln(f, string(jsonData)); err != nil {
	// 	log.Fatal(err)
	// }

	// json反序列化
	var stu2 Student
	jsonData, err := ioutil.ReadFile("day3/structDemo/info.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(jsonData, &stu2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(stu2)
}
