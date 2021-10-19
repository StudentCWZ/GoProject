package main

import "fmt"

// 演示 Golang 中 + 的使用
func main()  {
	var i = 1
	var j = 1
	var r = i + j // 做加法运算
	fmt.Println("r =", r)

	var str1 = "hello "
	var str2 = "world"
	var res = str1 + str2 // 做拼接操作
	fmt.Println("res =", res)
}

