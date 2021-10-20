package main

import "fmt"

func main() {
	/*
		逻辑运算符
	1. 用于连接多个条件(一般来讲就是关系表达式)，最终的结果也是一个 bool 值
	2. &&: 逻辑与运算符，如果两边的操作数都是 True，则为 True，否则为 False
	3. ||: 逻辑或运算符，如果两边的操作数有一个 True，则为 True，否则为 False
	4. !: 逻辑非运算符，如果条件为 True，则逻辑为 False，否则为 True
	*/
	// 演示逻辑运算符的使用 &&
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}

	// 演示逻辑运算符的使用 ||
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}
	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	// 演示逻辑运算符的使用 !
	if age > 30 {
		fmt.Println("ok5")
	}
	if !(age > 30) {
		fmt.Println("ok6")
	}

}
