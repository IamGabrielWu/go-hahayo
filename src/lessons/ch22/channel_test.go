package channel_test

import (
	"fmt"
	"sync"
	"testing"
)

// # channel 的关闭
// * 向关闭的channel 发送数据，会导致panic
// * v,ok <-ch ; ok 为bool 值， true 表示正常接受， false 表示通道关闭
// * 所有的channel 接收者都会在channel 关闭时， 立刻从阻塞等待中返回且上述ok值为false。 这个广播机制常被利用， 进行向多个订阅者同时发送信号， 如: 退出信号。

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		ch <- 11
		wg.Done()
	}()
}
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 11; i++ {
			if data, ok := <-ch; ok {
				fmt.Println(data, ok) // 通道接收， ok 返回true
			} else {
				fmt.Println(data, ok) // 通道关闭， ok 返回false
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	// wg.Add(1)
	// dataReceiver(ch, &wg)
	wg.Wait()
}
