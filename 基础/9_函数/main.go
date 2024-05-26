package main

import "fmt"

func main() {
	a := adder()
	fmt.Println(a(1)) // 输出: 1
	fmt.Println(a(2)) // 输出: 3，因为 sum 变量在闭包中保持状态
	fmt.Println(a(3)) // 输出: 6

	b := adder()
	// 调用闭包 b，它有自己的 sum 变量，与 a 无关
	fmt.Println(b(1)) // 输出: 1
	fmt.Println(b(2)) // 输出: 3
}

func adder() func(int) int {
	sum := 0 // 为什么这个sum能保持状态呢？go 为了实现闭包将栈变量逃逸到了堆上。
	return func(x int) int {
		sum += x
		return sum
	}
}
