package main

import "fmt"

func main()  {
	/*
		算术运算符的细节说明：
	1. 对于除号 "/"，它的整数除和小数除是有区别的，整数之间做除法时，只保留整数部分而舍弃小数部分，例如：x := 19 / 5，结果时 3
	2. 当对一个数取模时，可以等价 a % b = a - a / b * b，这样我们可以看到取模的一个本质运算
	3. Golang 的自增自减只能当做一个独立语言使用，不能这样使用 b := a++
	4. Golang 的 ++ 和 -- 只能写在变量的后面，不能写在变量的前面
	5. Golang 设计者去掉 c/java 中的自增自减的容易混淆的写法，让 Golang 更加简洁，统一。
	*/
	// 在 Golang 中，++ 和 -- 只能独立使用
	var i int = 8
	var a int
	// a = i++ 报错
	i++
	a = i
	fmt.Println(a)
	/* 报错
	if i++ > 0 {
		fmt.Println("ok")
	}
	*/
}
