package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	// b = a //无法转换，不支持隐式类型转换
	b = int64(a) // Go 只支持显式的类型转换
	t.Log(a, b)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a // a 的内存地址
	// aPtr = aPtr + 1 // 支持指针类型， 但不支持指针运算， 用指针访问连续的字符空间或者数组
	t.Log(a, aPtr)
	// t.Log("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string // string  默认是空字符串
	t.Log("*" + s + "*")
	t.Log(len(s))
}
