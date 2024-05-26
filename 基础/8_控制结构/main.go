package main

import "fmt"

func main() {
	//ifElse()
	//switchFunc()
	//forFunc()
	gotoFunc()
}

func ifElse() {
	var first int = 10
	var cond int

	// if - else if - else
	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {

		fmt.Printf("first is between 0 and 5\n")
	} else {

		fmt.Printf("first is 5 or greater\n")
	}

	// 初始化并 判断
	if cond = 5; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	} else {

		fmt.Printf("cond is not greater than 10\n")
	}
}

func switchFunc() {
	// 变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数
	// 一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块，也就是说您不需要特别使用 break 语句来表示结束。
	// 如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字来达到目的。
	//switch var1 {
	//case val1:
	//	...
	//case val2:
	//	...
	//default:
	//	...
	//}

	// 类似php的switch
	//var num1 int = 100
	//
	//switch num1 {
	//case 98, 99:
	//	fmt.Println("It's equal to 98")
	//case 100:
	//	fmt.Println("It's equal to 100")
	//default:
	//	fmt.Println("It's not equal to 98 or 100")
	//}

	// 另一种形式的switch ,类似if-else
	var num1 int = 7

	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}
}

func forFunc() {
	// 1. 基本的迭代
	//for i := 0; i < 5; i++ {
	//	fmt.Printf("This is the %d iteration\n", i)
	//}

	// 2. 基于条件判断的迭代
	//var i int = 5
	//for i >= 0 {
	//	i = i - 1
	//	fmt.Printf("The variable i is now: %d\n", i)
	//}

	// 3.无限循环
	//for {
	//
	//}

	// 4. for-range

}

func gotoFunc() {
	// 使用标签和 goto 语句是不被鼓励的：它们会很快导致非常糟糕的程序设计，而且总有更加可读的替代方案来实现相同的需求。
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
