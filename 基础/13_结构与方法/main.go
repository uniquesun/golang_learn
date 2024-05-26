package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
//type Address struct {
//	Id       uint
//	country  string
//	province string
//	city     string
//	address  string
//}
//
//type User struct {
//	Username string `json:"username"` // tags
//	Age      uint   `json:"age"`
//	Password string `json:"password"`
//	Avatar   string `json:"avatar"`
//	IsMan    bool   `json:"isMan"`
//	Address
//}
//
//// NewUser 工厂方法创建
//func NewUser(username string, age uint, Password string, avatar string, isMan bool) *User {
//	u := &User{Username: username, Age: age, Password: Password, Avatar: avatar, IsMan: isMan}
//	return u
//	//return &user{
//	//	username, age, Password, avatar, isMan,
//	//}
//}
//
//// 内嵌结构体
//type innerS struct {
//	in1 int
//	in2 int
//}
//
//type outerS struct {
//	b      int
//	c      float32
//	int    // anonymous field
//	innerS //anonymous field
//}
//
//func main() {
//
//	// 声明结构体
//	//var user = new(User)
//	//user.Username = "olaf"
//	//user.Age = 23
//	//fmt.Println(user)
//
//	// 工厂方法
//	//var user1 = NewUser("olaf", 29, "12313", "qwewq", true)
//	//fmt.Println(user1)
//
//	// tags
//	//refTag(user1)
//
//	//匿名字段 & 内嵌结构体
//	outer := new(outerS)
//	outer.b = 6
//	outer.c = 7.5
//	outer.int = 60
//	outer.in1 = 5
//	outer.in2 = 10
//
//	fmt.Printf("outer.b is: %d\n", outer.b)
//	fmt.Printf("outer.c is: %f\n", outer.c)
//	fmt.Printf("outer.int is: %d\n", outer.int)
//	fmt.Printf("outer.in1 is: %d\n", outer.in1)
//	fmt.Printf("outer.in2 is: %d\n", outer.in2)
//
//	// 结构体字面量
//	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
//	fmt.Println("outer2 is:", outer2)
//}
//
//func refTag(user *User) {
//	userType := reflect.TypeOf(user)
//	userValue := reflect.ValueOf(user)
//	length := userValue.Elem().NumField()
//
//	for i := 0; i < length; i++ {
//		field := userType.Elem().Field(i)
//		jsonName := field.Tag.Get("json")
//		fieldName := field.Name
//		fmt.Println(fieldName, jsonName)
//	}
//
//	//
//	//fmt.Printf("%v\n", ixField.Tag)
//}
