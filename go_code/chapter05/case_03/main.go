package main

import "fmt"

func main()  {
	/*
			双分支
	1. 基本语法
		if 条件表达式 {
		执行代码块1
	} else {
		执行代码块2
	}
	2. 说明：
		当条件表达式成立，即执行代码块1，否则执行代码块2。{} 也是必须有的
	*/
	// 应用案例
	/*
	编写一个程序，可以输入人的年龄，如果该同志的年龄大于 18 岁，则输出"你的年龄大于18，要对自己的行为负责！"。
	否则，输出"你的年龄不大这次放过你了"。
	*/
	var age int
	fmt.Println("请输入年龄：")
	_, _ = fmt.Scanln(&age)

	if age > 10 {
		fmt.Println("你的年龄大于18，要对自己的行为负责！")
	} else {
		fmt.Println("你的年龄不大这次放过你了")
	}

}