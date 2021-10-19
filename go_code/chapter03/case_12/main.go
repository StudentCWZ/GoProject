package main

import "fmt"

func main()  {
	/*
		基本数据类型的默认值
	1. 在 go 中，数据类型都有一个默认值，当程序员没有赋值时，就会保留默认值，在 go 中，默认值又叫零值。
	*/
	var a int // 0
	var b float32 // 0
	var c float64 // 0
	var isMarried bool // false
	var name string // ""
	fmt.Printf("a = %d, b=%f, c=%f, isMarried=%v, name=%v \n", a, b, c ,isMarried, name)
}

