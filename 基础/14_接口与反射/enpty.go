package main

import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

func main() {
	// 接口变量 val 被依次赋予一个 int，string 和 Person 实例的值，然后使用 type-switch 来测试它的实际类型。
	//每个 interface {} 变量在内存中占据两个字长：一个用来存储它包含的类型，另一个用来存储它包含的数据或者指向数据的指针。
	var val Any
	val = 5
	//fmt.Printf("val has the value: %v\n", val)
	val = str
	//fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	//fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}
