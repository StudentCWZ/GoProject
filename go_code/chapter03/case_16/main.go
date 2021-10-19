package main

import (
	"fmt"
	"strconv"
)

func main()  {
	/*
		基本数据类型和 string 转换
	1. 在程序开发中，我们经常将基本数据类型转成 string，或者将 string 转成基本数据类型
	*/

	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string // 空的 str

	// 方式一：
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type is %T str = %q \n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type is %T str = %q \n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type is %T str = %q \n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type is %T str = %q \n", str, str)

	// 方式二：使用 strconv 包的函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type is %T str = %q \n", str, str)

	// 说明：'f' 格式，10：表示小数位保留 10 位，64：表示这个小数是 float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type is %T str = %q \n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type is %T str = %q \n", str, str)

	// strconv 包中有一个函数 Itoa
	var num5 int = 4567
	str = strconv.Itoa(num5)
	fmt.Printf("str type is %T str = %q \n", str, str)
}


