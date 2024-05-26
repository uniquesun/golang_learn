package main

import "fmt"

func main() {
	// 字符串转字节数组
	name := "olaf"
	byteArr := []byte(name)
	fmt.Println(byteArr)
	// 字节数组转字符串
	fmt.Println(string(byteArr))
}
