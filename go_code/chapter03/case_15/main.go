package main

import "fmt"

func main()  {
	// 课堂练习
	var n1 int32 = 12
	var n3 int8
	// var n4 int8
	n3 = int8(n1) + 127 // 编译通过，按溢出处理
	// n4 = int8(n1) + 128 编译无法通过，128 超出 int8 取值范围
	fmt.Println("n3 =", n3)
}

