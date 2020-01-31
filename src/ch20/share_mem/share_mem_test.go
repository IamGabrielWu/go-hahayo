package share_mem

import (
	"sync"
	"testing"
	"time"
)

// # 共享内存并发机制
// 1. package sync: Mutex, RWLock

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() { // 每个协程，执行完之后最后释放锁
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// # WaitGroup
// var wg sync.WaitGroup
// for i:=0;i<5000;i++{
//   wg.Add(1)
//   go func ()  {
//     defer func ()  {
//       wg.Done() // 等到所有的wait 都执行完之后再执行wg.wait()之后的
//     }()
//     ...
//   }()
// }
// wg.Wait()

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() { // 每个协程，执行完之后最后释放锁
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}
