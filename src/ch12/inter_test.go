package inter_test

import (
	"fmt"
	"testing"
	"time"
)

// Duck Type 式接口实现
// 接口定义
// type Programmer interface {
//   WriteHelloWorld() Code
// }
// 接口实现： Go 接口的实现， 不依赖于接口的定义， 采用的是Duck Type 的方式， Duck Type 意思就是定义一个和接口内方法相同的方法
// type GoProgrammer struct {
//
// }
// Go 绑定方法不是在结构上，而是在指针上。
// func (p *GoProgrammer) WriteHelloWorld() Code {
//   return "fmt.Println(\"Hello World\")"
// }

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	p := new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

// Go 接口
// 与其他语言差异
// 1. 接口为非侵入性， 实现不依赖于接口定义
// 2. 所有接口的定义可以包含在使用者包内
// 接口变量
// var prog Coder=&GoProgrammer{}
// type GoProgrammer struct{}
// &GoProgrammer{}
// 自定义类型
// 1. type IntConvertionFn func(n int)  int
// 2. type MyPoint int
type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
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
	tsSF := timeSpent(slowFunc)
	tsSF(10)
}
