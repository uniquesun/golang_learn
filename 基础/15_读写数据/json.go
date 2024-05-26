package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	pa := &Person{"olaf", 12}
	// 序列化
	js, _ := json.Marshal(pa)
	fmt.Println(string(js))

	jsonData := []byte(`{"name":"Jane Smith","age":25}`)
	// 反序列化到任意数据
	var f interface{}
	json.Unmarshal(jsonData, &f)

	// 反序列化到结构体
	var p Person
	json.Unmarshal(jsonData, &p)
	fmt.Println(p)
}
