package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 函数是一等公民
// 与其他语言的区别
// 1.可以有多个返回值
// 2. 所有参数都是有值的传递， slice, map, channel 会传引用的错觉
// 3. 函数可以作为变量的值
// 4. 函数可以作为参数和返回值
func returnMultiVal() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 函数式编程， 将函数传入， 然后函数内定义要执行的函数， 然后返回函数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}
func slowFunc(op int) int {
	time.Sleep(time.Second * 2)
	return op
}
func TestFn(t *testing.T) {
	a, b := returnMultiVal()
	t.Log(a, b)
	a, _ = returnMultiVal() // _ 下划线， 忽略其中一个返回值
	t.Log(a)
	tsSF := timeSpent(slowFunc)
	tsSF(10)
}
