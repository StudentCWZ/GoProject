package main

import "fmt"

func main()  {
	/*
		程序流程控制介绍
	1. 在程序中，程序运行的流程控制决定程序是如何执行的，是我们必须掌握的，主要有三大流程控制语句
		(1) 顺序控制
		(2) 分支控制
		(3) 循环控制
	2. 顺序控制
		程序从上到下逐行地执行，中间没有任何判断和跳转。
	3. 顺序控制的举例
	*/
	// 顺序控制的举例: 代码中没有判断、没有跳转，程序按照默认的流程执行
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("%d 个星期零 %d 天 \n", week, day)

	// 第二题
	var huashi float32 = 134.2
	var sheshi float32  = 5.0 / 9 * (huashi - 100)
	fmt.Printf("%v 对应的摄氏温度= %v",huashi, sheshi)

	// Golang 中定义变量时采用合法的前向引用
	// 正确形式
	var num1 int = 10 // 声明 num1
	var num2 int = num1 + 20 // 使用 num1
	fmt.Println(num2)

	// 错误形式
	/*
	func main() {
		var num2 int  = 10 // 使用 num1
		var num1 int = 10 // 声明 num1
		fmt.Println(num2)
	}
	*/
}