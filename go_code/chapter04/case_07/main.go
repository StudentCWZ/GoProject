package main

import "fmt"

func test() int {
	return 90
}


func main()  {
	/*
		赋值运算符基本介绍
		1. 赋值运算符就是将某个运算后的值，赋给指定变量。
	*/
	// 赋值运算符
	//var i int
	//i = 10 // 基本赋值

	// 有两个变量，a 和 b，要求将其进行交换，最终打印结果
	// a = 9, b = 2 --> a = 2, b = 9
	a := 9
	b := 2
	fmt.Printf("交换前的情况是 a = %v, b = %v \n", a, b)
	// 定义一个临时变量
	t := a
	a = b
	b = t
	fmt.Printf("交换后的情况是 a = %v, b = %v \n", a, b)

	// 复合赋值的操作
	a += 17 // 等价 a = a + 17
	fmt.Println("a =", a)

	/*
		赋值运算符的特点
	1. 运算顺序从右往左
	2. 赋值运算符的左边只能是变量，右边可以是变量、表达式、常量值
	3. 复合赋值运算符等价于下面的效果：a += 3 等价于 a = a + 3
	*/
	var c int
	c = a + 3 // 赋值运算执行顺序是从右向左
	fmt.Println("c =", c)

	// 赋值运算符的左边只能是变量，右边可以是变量、表达式、常量值
	var d int
	d = a
	d = 8 + 2 * 8 // = 的右边是一个表达式
	d = test()
	d = 890 // 常量
	fmt.Println("d =", d)
}
