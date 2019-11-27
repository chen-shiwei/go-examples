package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 字符串分词
	s1 := "hello world!"
	for key, value := range strings.Fields(s1) {
		fmt.Printf("%v=>%v\n", key, value)
	}
	// 0=>hello
	// 1=>world!

	// 自定义去除字符
	words := strings.FieldsFunc(s1, func(r rune) bool {
		// 去除特殊字符
		return !unicode.IsLetter(r)
	})
	for key, value := range words {
		fmt.Printf("%v=>%v\n", key, value)
	}
	// 0=>hello
	// 1=>world

}
