package main

import "fmt"

// 演示 Golang 中小数类型使用
func main()  {
	var price float32 = 89.123
	fmt.Println("price =", price)
	// 关于浮点数在机器中存放形式：浮点数=符号位+指数位+尾数位
	var num1 float32 = -0.00089
	var num2 float64 = -7809656.09
	fmt.Println("num1 =", num1, "num2 =", num2)
	// 尾数部分可能丢失，造成精度损失。
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3 =", num3, "num4 =", num4)
	// 说明：float64 比 float32 要准确
	// 如果我们要保存精度高的数，应该选用 float64 位
	// 浮点型的存储分为三部分：符号位+指数位，在存储过程中，精度会丢失
	// Golang 的浮点类型有固定的范围和字段长度，不受具体 OS(操作系统) 的影响
	// Golang 的浮点型默认声明为 float64 类型
	var num5 = 1.1
	fmt.Printf("num5 的数据类型是 %T \n", num5)
	// 十进制数形式
	num6 := 5.12
	num7 := .123 // => 0.123
	fmt.Println("num6 =", num6, "num7 =", num7)
	// 科学计数法形式
	num8 := 5.1234e2 // 5.1234 * 10 的 2 次方
	num9 := 5.1234E2 // 5.1234 * 10 的 2 次方
	num10 := -5.1234E-2 // -0.051234
	fmt.Println("num8 = ", num8, "num9 =", num9, "num10 =", num10)

}

