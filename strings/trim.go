package main

import (
	"fmt"
	"strings"
)

func main() {

	var s = "1 a 1"
	trimS := strings.Trim(s, "1")
	fmt.Printf("Trim 前:%s 长度:%d \n后:%s 长度:%d \n", s, len(s), trimS, len(trimS)) // 前:1 a 1 长度:5 // 后: a  长度:3

	trimLeftS := strings.TrimLeft(s, "1")
	fmt.Printf("TrimLeft 前:%s 长度:%d \n后:%s 长度:%d \n", s, len(s), trimLeftS, len(trimLeftS)) // TrimLeft 前:1 a 1 长度:5  // 后: a 1 长度:4

	trimRightS := strings.TrimRight(s, "1")
	fmt.Printf("TrimRight 前:%s 长度:%d \n后:%s 长度:%d \n", s, len(s), trimRightS, len(trimRightS)) // TrimRight 前:1 a 1 长度:5  // 后:1 a  长度:4

	trimSpaceS := strings.TrimSpace(" abc ")
	fmt.Printf("TrimSpace 前:%s 长度:%d \n后:%s 长度:%d \n", " abc ", len(" abc "), trimSpaceS, len(trimSpaceS)) // TrimRight 前:1 a 1 长度:5  // 后:1 a  长度:4

}
