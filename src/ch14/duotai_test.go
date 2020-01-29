package duotai_test

import (
	"fmt"
	"testing"
)

type Code string // 自定义类型， string 类型
type Programmer interface {
	WriteHelloWorld() Code
}
type GoProgrammer struct{}
type JavaProgrammer struct{}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Go - Hello World\")"
}
func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Java - Hello World\")"
}
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}

// 空接口与断言
// 1. 空接口可以表示任何类型， 相当于Java 里的Object 类型
// 2. 通过断言来将空接口转换成制定的类型
// v, ok :=p.(int) //ok=true 时为转换成功

// p.()， （）表示断言， 括号内为断言内容，为什么类型
func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("string", s)
	// 	return
	// }
	// fmt.Println("Unknown Type")
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown Type")
	}
}

func TestEmptyINterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("hello world")
}

// Go 接口最佳实践
// 倾向于使用小的接口定义，很多接口只包含一个方法
// type Reader interface {
//   Read(p []byte) (n int, err error)
// }
// type Writer interface {
//   Write(p []byte)(n int, err error)
// }
// 较大的接口定义，可以由多个小接口定义组合而成
// type ReadWriter interface {
//   Reader
//   Writer
// }
// 只依赖于必要功能的最小接口
// func StoreData(reader Reader) error {
//
// }
