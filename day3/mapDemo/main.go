// map定义

package main

import "fmt"

func main() {
	// 定义一
	info := make(map[string]string)
	info["Tom"] = "1827"
	info["jackon"] = "8271"

	for k, v := range info {
		fmt.Println(k, v)
	}

	// 定义二
	ages := map[string]int{"Rose": 18, "Jack": 23}
	for k, v := range ages {
		fmt.Println(k, v)
	}

	// 获取元素, 添加元素
	if element, ok := ages["Rose"]; ok {
		fmt.Println(element)
	} else {
		fmt.Println("not found")
	}

	// 不重复添key加新元素
	if _, ok := ages["Jackon"]; ok {
		fmt.Println("key存在!")
	} else {
		ages["Jaskon"] = 59
	}
	fmt.Println(ages)

	// 删除元素
	if _, ok := ages["Rose"]; ok {
		delete(ages, "Rose")
	} else {
		fmt.Println("not found")
	}
	fmt.Println(ages)
}
