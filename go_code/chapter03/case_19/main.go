package main

import "fmt"

func main()  {
	/*
		课堂练习：
	1. 写一个程序，获取一个 int 变量 num 的地址，并显示到终端
	2. 将 num 的地址赋给指针 ptr，并通过 ptr 去修改 num 值
	*/
	var num int = 9
	fmt.Printf("num address = %v \n", &num)
	var ptr *int = &num
	*ptr = 10
	fmt.Println("num =", num)
}

