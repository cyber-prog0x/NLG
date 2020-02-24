package main

import (
	"fmt"
)

func floatDemo() {
	//math.MaxFloat32 // 32位浮点数最大
	f1 := 1.23456 // 默认Go语言中的小数都是float64类型
	fmt.Printf("%T \n", f1)

	f2 := float32(1.23456)
	fmt.Printf("%T \n", f2) // 显示声明float32类型name

	//f1 = f2 // float32 type can not assign to float64 type

}

func complexdemo() {
	//var c1 complex64
}

func booldemo() {
	// Go 语言中不允许将整型强制类型转换为布尔类型

	b1 := true
	var b2 bool

	fmt.Printf("%T\n", b1)
	fmt.Printf("%T\n", b2)

	fmt.Printf("%v %v", b1, b2)
}

func main() {
	var a int = 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a)

	var b int = 077 // 八进制
	fmt.Printf("%o \n", b)

	var c int = 0xff
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)

	// 查看变量类型
	fmt.Printf("%T \n", c)

	i4 := int8(9) // 强制定义为 int8类型 否则默认则是int类型
	fmt.Printf("%T \n", i4)

	// float type
	floatDemo()
	booldemo()

}
