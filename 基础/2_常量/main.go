package main

import "fmt"

func main() {
	// 定义常量
	const PI = 3.14159
	// 枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	// 可以把iota理解为每行的index，后面不设置的话就依次 等于前面的公式（iota * 10）
	const (
		Sunday = 70
		Monday
		Tuesday = iota * 10
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(PI)
}
