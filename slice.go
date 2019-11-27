package main

import "fmt"

func main() {
	isNil()
	sliceCopy()
}

func isNil() {

	var s1 []int
	fmt.Printf("s1 == nil %v \n", s1 == nil) // s1 == nil true

	s2 := []int{}
	fmt.Printf("s2 == nil %v \n", s2 == nil) //  s2 == nil false

	s3 := make([]int, 0)
	fmt.Printf("s3 == nil %v \n", s3 == nil) //  s3 == nil false

}

func sliceCopy() {

	// 浅拷贝
	s1 := make([]int, 5, 5)
	s2 := s1
	s1[1] = 1
	fmt.Printf("s1：%v ，s2：%v \n", s1, s2) // s1：[0 1 0 0 0] ，s2：[0 1 0 0 0]
	// 深拷贝
	s3 := make([]int, 5, 5)
	s4 := make([]int, 5, 5)
	copy(s4, s3)
	s3[1] = 1
	fmt.Printf("s3：%v ，s4：%v \n", s3, s4) // s3：[0 1 0 0 0] ，s4：[0 0 0 0 0]

	// 深拷贝
	s5 := []int{1, 2, 3, 4, 5}
	s6 := s5
	s5 = append(s5, 6)
	fmt.Printf("s5：%v ，s6：%v \n", s5, s6) // s5：[1 1 3 4 5 6] ，s6：[1 1 3 4 5]
	s5[2] = 0
	fmt.Printf("s5：%v ，s6：%v \n", s5, s6) // s5：[1 2 0 4 5 6] ，s6：[1 2 3 4 5]
	s6[2] = 8
	fmt.Printf("s5：%v ，s6：%v \n", s5, s6) // s5：[1 2 0 4 5 6] ，s6：[1 2 8 4 5]

}
