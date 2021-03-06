package main

/*
	值类型和引用类型
- 值类型和引用类型的说明
1. 值类型：基本数据类型 int 系列，float 系列，bool，string，数组和结构体 struct
2. 引用类型：指针、slice 切片、map、管道 channel、interface 等都是引用类型
- 值类型和引用类型使用特点
1. 值类型：变量直接存储值，内存通常在栈中分配
2. 应用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据(值)，内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就
成为一个垃圾，由 GC 回收
*/
