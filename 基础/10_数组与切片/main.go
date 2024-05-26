package main

import "fmt"

func main() {
	sliceFunc()
}

func arrayFunc() {
	// 声明数组,初始值 [0 0 0 0 0]
	var arr1 [5]int // arr1是指针 *[5]int
	fmt.Println(arr1)

	// 数组是一种 值类型, 也可以通过new来创建
	var arr2 = new([5]int) // arr2 是 值[5]int
	fmt.Println(arr2)

	// 简写声明
	//var arrAge = [5]int{18, 20, 15, 22, 16}
	//var arrLazy = [...]int{5, 6, 7, 8, 22} // 让编译器判断长度
	//var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}  // 声明index3 = Chris

	// 数组赋值
	arr1 = [5]int{12, 32, 43, 54, 2}

	// 遍历 for-i
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	// 遍历 for-range
	for _, v := range arr1 {
		fmt.Println(v)
	}

}

func sliceFunc() {
	// 申明切片 一个切片在未初始化之前默认为 nil，长度为 0。
	var s1 []string
	fmt.Println(s1)

	// make函数

	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	// 遍历
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("切片长度 %d\n", len(slice1))
	fmt.Printf("切片容量 %d\n", cap(slice1))

	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

}
