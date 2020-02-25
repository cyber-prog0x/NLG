package main

import (
	"fmt"
	"strings"
)

//Go语言中字符串是用双引号包裹
//Go语言中单引号包裹的是字符

func main() {
	/* s := "Hello 世界"*/

	//c1 := 'h'
	//c2 := '1'
	//c3 := '世'

	// 字节 1byte == 8bit

	// 1个字符 'A' = 一个字节
	// 1个UTF-8编码的中文字符一般占用 3个字节

	// 字符串转义
	// \r 回车(返回行首)
	// \n 换行符
	// \t 制表符
	// \' 单引号
	// \" 双引号
	// \\ 反斜杠

	fmt.Println("C:\\Windows\\System32\\etc\\hosts")

	// 多行字符串

	s := `
		我们都是世界的公民
		愿世界和平
		
	`

	fmt.Println(s)

	// 字符串操作
	fmt.Println(len(s))

	name := "dream"
	describe := "come true"
	fmt.Println(name + describe)

	s3 := fmt.Sprintf("%s%s", name, describe)
	fmt.Println(s3)

	ret := strings.Split(s3, "co")
	fmt.Println(ret)

	// strings.Contains
	fmt.Println(strings.Contains(s3, "world"))
	// strings.HasPrefix
	fmt.Println(strings.HasPrefix(s3, "dr"))
	// Strings.HasSuffix
	fmt.Println(strings.HasSuffix(s3, "ur"))

	// index op
	s4 := "abcdefd"
	fmt.Println(strings.Index(s4, "d"))
	fmt.Println(strings.LastIndex(s4, "d"))

	// Join
	s5 := strings.Join(s4, " hello")
	fmt.Println(s5)

}
