package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		string 类型转基本数据类型
	1. 使用 strconv 包的函数
	*/
	// 字符串 --> 布尔型
	var str string = "true"
	var b bool
	// 说明：
	// 1. strconv.ParseBool(str)，函数返回两个值(value bool, err error)
	// 2. 因为我只想获取到 value bool，不想获取 err 所以我使用 _ 忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b = %v \n", b, b)

	// 字符串 --> 整型
	var str2 string = "1234590"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1 = %v \n", n1, n1)
	fmt.Printf("n2 type %T n2 = %v \n", n2, n2)

	// 字符串 --> 浮点型
	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1 = %v \n", f1, f1)

	// 注意事项：在将 string 类型转成基本数据类型时，要确保 string 类型能够转成有效的数据。
	// 比如：我们可以把 "123" 转成一个整数，但是不能把 "hello" 转成一个整数，如果这样做，Golang 直接将其转成 0 。
	var str4 string = "hello"
	var n3 int64 = 11
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type is %T n3 = %v \n", n3, n3)

	var b2 bool
	b2, _ = strconv.ParseBool(str4)
	fmt.Printf("b2 type %T b2 = %v \n", b2, b2)
}

