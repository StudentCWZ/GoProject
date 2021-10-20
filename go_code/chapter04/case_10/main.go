package main

import "fmt"

func main()  {
	/*
		其他运算符
	1. &: 返回变量存储地址
	2. *: 指针变量
	*/
	// 演示 & 和 * 的使用
	a := 100
	fmt.Println("a 的地址 =", &a)
	var ptr *int = &a
	fmt.Println("ptr 指向的值 =", *ptr)

	//
	var n int
	var i int = 10
	var j int = 12
	// 传统三元运算
	// n = i > j ? i:j
	if i > j {
		n = i
	} else{
		n = j
	}
	fmt.Println("n =", n) // 12
}
