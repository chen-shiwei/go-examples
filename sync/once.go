package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var one sync.Once
	// 只运行一次
	for {
		one.Do(func() {
			fmt.Println(time.Now().Format(time.RFC3339)) // only once 2019-11-27T15:44:03+08:00
		})
	}
}
