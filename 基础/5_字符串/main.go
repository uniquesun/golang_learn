package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Hello,世界"

	// 遍历字符串
	for _, v := range name {
		fmt.Println(string(v))
	}

	// 判断开头
	fmt.Println(strings.HasPrefix(name, "hello"))

}
