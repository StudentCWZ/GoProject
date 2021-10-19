package main

import "fmt"

func main()  {
	// Golang 的变量使用方式：
	// 第一种：指定变量类型，声明后若不赋值，使用默认值
	var i int
	fmt.Println("i =", i)
	// 第二种：根据值自定判定变量类型(类型推导)
	var num = 10
	fmt.Println("num =", num)
	// 第三种：省略 var，注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误
	// 下面的方式等价 var name string     name = "Tom" 两句
	name := "Tom"
	fmt.Println("name =", name)
}

