package main

import (
	"errors"
	"fmt"
)

// 定义错误
var errNotFound error = errors.New("Not found error")

func main() {

	fmt.Printf("error: %v", errNotFound)

	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!") // 恐慌
	fmt.Println("Ending the program")
}
