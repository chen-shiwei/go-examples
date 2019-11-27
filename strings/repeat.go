package main

import (
	"fmt"
	"strings"
)

func main() {

	// 重读字符串填充 类似于 php 函数 str_repeat(string,repeat)

	s := "abc"
	s1 := strings.Repeat(s, 3)

	fmt.Printf("s:%s s1:%s", s, s1) // out s:abc s1:abcabcabc
	return

}
