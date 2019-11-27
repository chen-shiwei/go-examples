package main

import (
	"fmt"
	"strings"
)

func main() {
	// 把数组转成字符串，并用指定字符连接
	var s = []string{"a", "b", "c"}
	s1 := strings.Join(s, ",")
	s2 := strings.Join(s, "-")

	fmt.Printf("s1:%s s2:%s", s1, s2) //outs1:a,b,c s2:a-b-c
}
