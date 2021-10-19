package main

import "fmt"

// 演示 Golang 中 string 类型使用
func main()  {
	/*
		字符串的基本介绍
	1. 字符串就是一串固定长度的字符连接起来的字符序列，Go 的字符串是由单个字节连接起来的。
	2. Go 语言的字符串的字节使用 UTF-8 编码标识文本
	*/
	// string 的基本使用
	var address string = "北京长城 110 hello world!"
	fmt.Println(address)
	/*
		字符串的注意事项
	1. 字符串一旦赋值，字符串就不能修改了，在 Go 中字符串是不可变的。
	*/
	// var str = "hello"
	// str[0] = 'a' 报错： 这里就不能去修改 str 内容，即 Go 中字符串是不可变的。
	/*
		字符串的表示形式
	1. 双引号：会识别转义字符
	2. 反引号：以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
	*/
	str2 := "abc\nabc"
	fmt.Println(str2)
	str3 := `abc\nabc`
	fmt.Println(str3)
	// 字符串拼接方式
	var str4 = "hello " + "world"
	str4 += " haha!"
	fmt.Println(str4)
	// 当一个拼接的操作很长，可以分行写，把加号放在上面
	var str5 = "hello " + "world" + "hello " + "world" + "hello " + "world" +
		"hello " + "world" + "hello " + "world"
	fmt.Println(str5)


}

