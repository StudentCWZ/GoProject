package main

import "fmt"

func main()  {
	/*
		Golang 基本数据类型的转换
	1. Golang 和 java / C 不同，Go 在不同类型的变量之间赋值时需要现式转换，也就是说 Golang 中数据类型不能自动转换。
	2. 基本语法:
		(1) 表达式 T(v) 将值 v 转换为类型 T
		(2) T：就是数据类型，比如 int32、int64 等
		(3) v：就是需要转换的变量
	*/
	var i int32 = 100
	// 希望将 i => float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	// var n3 int64 = i // 低精度 -> 高精度(部分语言可以，但是 Golang 不可以)
	var n3 int64 = int64(i)
	fmt.Printf("i = %v n1 = %v n2 = %v n3 = %v \n", i, n1, n2, n3)
	/*
		Golang 基本数据类型转换的细节说明
	1. Go 中，数据类型的转换可以是从表面范围小->范围大，也可以范围大->范围小
	2. 被转换的是变量存储的数据，变量本身的数据类型并没有变化！
	3. 在转换中，比如将 int64 转成 int8，编译时不会报错，只是转换的结果是按溢出处理，和我们希望结果不一样。
	*/
	// 被转换的是变量存储的数据，变量本身的数据类型并没有变化！
	fmt.Printf("i type is %T \n", i)
	// 在转换中，比如将 int64 转成 int8，编译时不会报错，只是转换的结果是按溢出处理，和我们希望结果不一样。
	var num1 int64 = 999999
	var num2 int8 = int8(num1) // 溢出处理
	fmt.Println("num2 =", num2)

}

