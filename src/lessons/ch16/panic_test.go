package panic_test

import (
	"errors"
	"fmt"
	"testing"
)

// panic
// 1. panic 用于不可恢复的错误
// 2. panic退出前会执行defer 指定内容
// panic vs os.Exit
// 1. os.Exit 退出时不会调用defer 指定函数
// 2. os.Exit  退出时不输出当前的调用栈信息
func TestPanicVsExit(t *testing.T) {
	defer func() {
		fmt.Println("finally!")
	}()
	fmt.Println("start")
	// os.Exit(-1)
	panic(errors.New("Something wrong!"))
	// fmt.Println("end")
}

// ### recover, 不中断程序执行， 捕捉错误，并且让程序继续执行
// defer func() {
// 	if err:=recover(); err!=nil {
// 		// 恢复错误
// 	}
// }()
func TestPanicVsExit2(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("start")
	// os.Exit(-1)
	panic(errors.New("Something wrong!"))
	// fmt.Println("end")
}

// # 当心! recover成为恶魔
// 1. 形成僵尸服务进程， 导致health check 失效
// 2. “let it crash” 往往是我们恢复不确定性错误的最好方法
