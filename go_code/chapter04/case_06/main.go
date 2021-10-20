package main

import "fmt"


// 声明一个函数(测试)
func test() bool {
	fmt.Println("test ...")
	return true
}

func main() {
	/*
		逻辑运算符注意事项和细节说明
	1. && 也叫短路与，如果第一个条件为 False，则第二个条件不会判断，最终结果为 False
	2. || 也叫短路或，如果第一个条件为 True，则第二个条件不会判断，最终结果为 True
	*/
	var i int = 10
	// 短路与
	if i > 9 && test() {
		fmt.Println("ok ...")
	}
	// 说明: 因为 i < 9 为 False，因此后面的 test() 不执行
	if i < 9 && test() {
		fmt.Println("ok1 ...")
	}
	// 短路或
	if i < 9 || test() {
		fmt.Println("ok2 ...")
	}
	// 说明：因为 i > 9 为 True，因为后面的 test() 就不执行
	if i > 9 || test() {
		fmt.Println("ok3 ...")
	}
}
