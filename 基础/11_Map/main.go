package main

import "fmt"

func main() {
	fun1()
}

func fun1() {
	// 声明,未初始化的 map 的值是 nil。
	var map1 map[string]string
	// 初始化
	map1 = map[string]string{"one": "one", "two": "two"}
	fmt.Println(map1)

	// make函数
	mapCreated := make(map[string]float32)
	mapCreated["score"] = 99.3
	fmt.Println(mapCreated)

	// 遍历
	for key, value := range mapCreated {
		fmt.Println(key)
		fmt.Println(value)
	}

	// 判断是否存在
	if _, ok := map1["score"]; ok {
		// ...
	}

	// 删除key
	delete(mapCreated, "score")
}
