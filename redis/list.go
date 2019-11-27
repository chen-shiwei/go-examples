package redis

import (
	"fmt"
)

func List() {

}

func LPush(key string, val ...interface{}) (rows int64, err error) {
	rows, err = cli.LPush(key, val).Result()
	if err != nil {
		err = fmt.Errorf("LPush err:%s", err.Error())
		return
	}
	return
}

func LPushX(key string, val string) (rows int64, err error) {
	rows, err = cli.LPushX(key, val).Result()
	if err != nil {
		err = fmt.Errorf("LPushX err:%s", err.Error())
		return
	}
	return
}

func TwoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum > target {
			j--
			continue
		}

		if sum < target {
			i++
			continue
		}

		return []int{i + 1, j + 1}

	}
	return []int{}
}

func PVowels() {
	return
}

var vowels = map[uint8]string{97: "a", 101: "e", 105: "i", 111: "o", 117: "u", 65: "A", 69: "E", 73: "I", 79: "O", 85: "U"}

func ReverseVowels(s string) string {
	i, j := 0, len(s)-1

	b := make([]byte, j+1)

	for i <= j {
		if _, ok := vowels[s[i]]; !ok {
			b[i] = s[i]
			i++
		} else if _, ok := vowels[s[j]]; !ok {
			b[j] = s[j]
			j--
		} else {
			b[i] = s[j]
			b[j] = s[i]
			i++
			j--
		}
	}
	return string(b)
}

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	i1 := m - 1
	i2 := n - 1
	mergeI := m + n - 1
	for i1 >= 0 || i2 >= 0 {
		if i1 < 0 {
			nums1[mergeI] = nums2[i2]
			i2--
			mergeI--
			continue
		}

		if i2 < 0 {
			nums1[mergeI] = nums1[i1]
			mergeI--
			i1--
			continue
		}

		if nums1[i1] > nums2[i2] {
			nums1[mergeI] = nums1[i1]
			mergeI--
			i1--
			continue
		}

		nums1[mergeI] = nums2[i2]
		mergeI--
		i2--
	}

	return nums1
}
