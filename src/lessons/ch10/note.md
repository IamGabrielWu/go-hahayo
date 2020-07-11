### defer
// 函数： 可变参数及延迟运行defer
1)
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clear resources")
	}() // 不会立即执行， 而是下面的语句执行完之后再执行。 相当于finally. 清理某些资源， 释放某些锁
	t.Log("Started")
	panic("Fatal error") // 相当于finally, defer 仍会执行
	// fmt.Println("test after panic") // panic 之后其他语句没法执行
}
2)
func Clear() {
	fmt.Println("Clear resoruces Again.")
}

func TestDeferAgain(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
}