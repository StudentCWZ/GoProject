package main

import "fmt"

func main()  {
	/*
		关系运算符/比较运算符
	1. 关系运算符的结果都是 bool 型，也就是要么是 true，要么是 false
	2. 关系表达式经常用在 if 结构的条件中或循环结构的条件中
	*/
	// 演示关系运算符的使用
	var n1 int = 9
	var n2 int = 8
	fmt.Println(n1 == n2) // false
	fmt.Println(n1 != n2) // true
	fmt.Println(n1 > n2)  // true
	fmt.Println(n1 >= n2) // true
	fmt.Println(n1 < n2)  // false
	fmt.Println(n1 <= n2) // false

	flag := n1 > n2
	fmt.Println("flag =", flag)
	/*
		关系运算符的细节说明
	1. 关系运算符的结果都是 bool 型，要么是 true，要么是 false
	2. 关系运算符组成的表达式，我们称为关系式：a > b
	3. 比较运算符 "==" 不能误写成 "="
	*/
}
