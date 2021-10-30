package main

import "fmt"

func main()  {
	/*
		分支控制 if-else:
	1. 让程序有选择的执行，分支控制有三种：
		(1) 单分支
		(2) 双分支
		(3) 多分支
	2. 单分支的基本语法
		if 条件表达式 {
		执行代码块
	}
		(1) 说明：当条件表达式为 true 时，就会执行 {} 的代码
		(2) 注意 {} 是必须有的，就算你只写一行代码
	*/
	// 单分支的应用案例
	// 编写一个程序，可以输出人的年龄，如果该同志的年龄大于18岁，则输出"你年龄大于 18，要对自己的行为负责"

	// 分析
	//1. 年龄 -> var age byte
	//2. 从控制台接收一个输入
	//3. if 判断
	//var age int
	//fmt.Println("请输入年龄：")
	//_, _ = fmt.Scanln(&age)
	//if age > 18 {
	//	fmt.Println("你年龄大于 18，要对自己的行为负责")
	//}
	// 细节说明：Golang 的 if 还有一个强大的地方就是条件判断语句里面允许声明一个变量，
	//这个变量的作用只能在该条件逻辑块内，其他地方就不起作用了。

	// Golang 支持在 if 中，直接定义一个变量，比如下面
	if age := 20; age > 18 {
		fmt.Println("你年龄大于 18，要对自己的行为负责")
	}
}