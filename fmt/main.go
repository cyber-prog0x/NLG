package main

import "fmt"

// fmt 占位符
func main() {
	var n = 100

	// show type
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	var s = "Hello 世界!"
	fmt.Printf("string: %s\n", s)
	fmt.Printf("string: %v\n", s)
	fmt.Printf("string: %#v\n", s)
}
