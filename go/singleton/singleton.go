// 利用sync.once实现一个单例
package main

import (
	"fmt"
	"sync"
)

// 均为小写，不可导出
var (
	once sync.Once
	s    singleton
)

// 定义单例时，一定要使用小写开头
// 使结构体不可导出
// 否则别的包就能直接使用&Singleton{}构建对象
type singleton struct {
	Val int
}

func GetSingleton() *singleton {
	// 仅创建一次
	once.Do(func() {
		s = singleton{Val: 100}
	})
	return &s
}

func main() {
	for i := 0; i < 10; i++ {
		s := GetSingleton()
		fmt.Printf("address: %p, value: %d\n", s, s.Val)
	}
}
