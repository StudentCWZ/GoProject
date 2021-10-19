package main

import "fmt"

// 演示 Golang 指针类型
func main() {
	/*
		Golang 指针
	1. 基本数据类型，变量存的就是值，也就是值类型
	2. 获取变量的地址，用 &，比如：var num int，获取 num 地址：&num
	3. 指针类型，变量存的是一个地址，这个地址指向的空间才是值，比如：var ptr *int = &num
	4. 获取指针类型所指向的值，使用：*，比如：var ptr *int，使用 *ptr 获取 p 指向的值
	*/
	// 基本数据类型的内存布局
	var i int = 10
	// i 地址
	fmt.Println("i 的地址 =", &i)

	// 下面的 var ptr *int = &i
	// 1. ptr 是一个指针变量
	// 2. ptr 的类型 *int
	//3. ptr 本身的值&i
	var ptr *int = &i
	fmt.Printf("ptr = %v \n", ptr)
	fmt.Printf("ptr 的地址 = %v \n", &ptr)
	fmt.Printf("ptr 指向的值 = %v \n", *ptr)
}

