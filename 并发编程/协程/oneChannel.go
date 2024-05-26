package main

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	//out <- 2 // 代码执行到这，没有协程接受，所以把这段代码放到 go f1(out)下面就行
	go f1(out)
	out <- 2
	time.Sleep(1)
}
