package main

import (
	"fmt"
	"unsafe"
)


// 演示 Golang 中 bool 类型使用
func main() {
	/*
		布尔类型
	1. 布尔类型也叫 bool 类型，bool 类型数据只允许取值 true 和 false
	2. bool 类型占 1 个字节
	3. bool 类型适用于逻辑运算，一般用于程序的流程控制
	*/
	var  b = false
	fmt.Println("b =", b)
	// 注意事项
	// 1. bool 类型占用存储空间时 1 个字节
	fmt.Println("b 的占用空间 =", unsafe.Sizeof(b))
	// 2. bool 类型只能取 true 或者 false
	// b = 1 报错
}

