package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

// # Thread 栈 vs Groutine
// 1. 创建时默认的stack 的大小
// * 创建时默认的stack 的大小
// * Groutine 的stack 初始化大小2K
// 2. 和KSE (Kernel Space Entity， 系统线程) 的对应关系
// * Java Thread 是1:1 , 当多线程时候，会造成上下文切换频繁
// * Groutine 是M:N， 多对多关系

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			time.Sleep(time.Second * 2)
		}(i)
	}
}
