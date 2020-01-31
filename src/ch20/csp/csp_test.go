package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}
func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// 声明一个channel ， 指定channel 内数据类型
func AsyncService() chan string {
	// retCh := make(chan string) //, 阻塞channel 意味着客户端调用get 把内容取完之后，程序才会继续往下执行。
	retCh := make(chan string, 1) // 这是buffer channel 。 不受客户端取还是不取的限制
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret // <- 往channel 里放入数据
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) // <- 往channel 里取出数据
}
