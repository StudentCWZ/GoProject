package main

import "fmt"

func main()  {
	// 课堂练习
	var a int = 300
	var b int = 400
	var ptr *int = &a
	*ptr = 100
	ptr = &b
	*ptr = 200
	fmt.Printf("a = %d, b = %d, *ptr = %d", a, b, *ptr)
	/*
		指针的使用细节
	1. 值类型，都有对应的指针类型，形式为 *数据类型，比如 int 的对应的指针就是 *int，float32 对应的指针类型就是 *float32
	2. 值类型包括：基本数据类型 int 系列，float 系列，bool，string，数组和结构体(struct)
	*/
}

