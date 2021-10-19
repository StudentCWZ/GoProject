package main

// 引用多个包
import (
	"fmt"
	"unsafe"
)

// 演示 Golang 中整数类型使用
func main()  {
	var i = 1
	fmt.Println("i =", i)

	// 测试一下 int8 范围
	// 其他的 int16，int32，int64 类推
	var j int8 = -128
	fmt.Println("j =", j)
	// var j int8 = -129 报错：超出范围
	var j1 int8 = 127
	fmt.Println("j1 =", j1)
	// var j1 int8 = 128 报错：超出范围

	// 测试一下 uint8 的范围，其他的 uint16，uint32，uint64 以此类推
	var k uint8 = 0
	fmt.Println("k =", k)
	// var k uint8 = -1 报错：超出范围
	var k1 uint8 = 255
	fmt.Println("k1 =", k1)
	// var k1 uint8 = 255 报错：超出范围
	// int，unit，rune，byte 的使用
	var a int = 8900
	fmt.Println("a =", a)
	var b uint  = 1
	fmt.Println("b =", b)
	var c byte = 0 // 范围 0～255，等价于 uint8
	fmt.Println("c =", c)
	var d rune = 7 // 等价于 int32
	fmt.Println("d =", d)

	// 整型的使用细节
	var n1 = 100 // n1 是什么类型
	// 这里我们给大家介绍一下如何查看某个变量的数据类型
	// fmt.Printf() 可以用做格式化输出
	fmt.Printf("n1 的类型 %T \n", n1)
	// 查看某个变量的占用字节大小和数据类型(使用较多)
	var n2 int64 = 10
	// unsafe.Sizeof(n2) 是 unsafe 包的一个函数，可以返回 n1 变量占用的字节数
	fmt.Printf("n2 的类型 %T n2 占用的字节数 %d", n2, unsafe.Sizeof(n2))
	// Golang 程序中整型变量在使用时，遵守保小不保大的原则。
	// 即：在保证程序正确运行下，尽量使用占用空间小的数据类型
	// var age byte = 90
	// bit: 计算机中的最小存储单位，byte：计算机中基本存储单元。1 byte = 8 bit
}

