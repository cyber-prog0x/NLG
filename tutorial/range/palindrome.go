package main

import "fmt"

func main() {
	s := "你好好你"
	// 需要把字符放到一个rune的切片里面

	r := make([]rune, 0, len(s))

	for _, c := range s {
		r = append(r, c)
	}

	fmt.Println(len(r))

	for i := 0; i < len(r)/2; i++ {
		// ss[0] s[len(s)-1]

		if r[i] != r[len(r)-1-i] {
			fmt.Println("It's not palindrome")
			return

		}
	}

	fmt.Println("It's palindrome")
}
