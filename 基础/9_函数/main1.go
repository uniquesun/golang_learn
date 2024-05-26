package main

import "fmt"

func main() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
	numx2, _ := getX2AndX3_2(1)
	fmt.Println(numx2)
	Greeting("hello", "olaf", "long", "jie")
	a()
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

// MultiPly3Nums 普通函数
func MultiPly3Nums(a, b, c int) int {
	return a * b * c
}

// 命名返回值
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

// Greeting 变长参数
func Greeting(prefix string, who ...string) {
	fmt.Println(prefix)
	fmt.Println(who) // who 是一个 slice 类型的变量
}

// defer
func a() {
	i := 0
	defer fmt.Println(i) // 0
	i++
	return
}

// 递归
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

// 闭包
