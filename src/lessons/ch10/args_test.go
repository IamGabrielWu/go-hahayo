package args_test

import (
	"fmt"
	"testing"
)

// 函数： 可变参数及延迟运行defer
// func sum(ops ...int) int {
//   s :=0
//   for _, op: range ops {
//     s+=op
//   }
//   return s
// }
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5, 6, 7, 8))
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clear resources")
	}() // 不会立即执行， 而是下面的语句执行完之后再执行。 相当于finally. 清理某些资源， 释放某些锁
	t.Log("Started")
	panic("Fatal error") // 相当于finally, defer 仍会执行
	// fmt.Println("test after panic") // panic 之后其他语句没法执行
}

func Clear() {
	fmt.Println("Clear resoruces Again.")
}

func TestDeferAgain(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
}
