package main

import "fmt"

// Person 封装
type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) SetFirstName(firstName string) {
	fmt.Println("Person 执行")
	p.FirstName = firstName
}
func (p *Person) GetFirstName() string {
	return p.FirstName
}

type Student struct {
	Person
	Score     float32
	FirstName string
}

func (s *Student) SetFirstName(fistName string) {
	fmt.Println("student 执行")
	s.FirstName = fistName
}

func main() {
	//p1 := new(Person)
	//p1.SetFirstName("olaf")
	//fmt.Println(p1.GetFirstName())

	s1 := &Student{
		Person:    Person{"olaf", "loong"},
		Score:     100,
		FirstName: "jie",
	}

	// 获取同名的属性
	fmt.Println(s1.FirstName) // jie
	// 获取嵌入结构体属性
	fmt.Println(s1.Person.FirstName) // olaf

	// 执行方法
	s1.SetFirstName("olaf student")
	// 执行嵌入结构体方法
	s1.Person.SetFirstName("xixi")
	//fmt.Println(s1.GetFirstName())
}
