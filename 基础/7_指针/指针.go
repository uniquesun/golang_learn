package main

import "fmt"

func main() {
	var i1 = 5
	var intP *int // 声明指针类型
	intP = &i1
	// 获取地址
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
	fmt.Println(intP)

	// 修改地址值
	*intP = 10
	fmt.Println(*intP)
	fmt.Println(i1)

}
