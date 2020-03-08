package series

import "fmt"

// # package
// 1. 基本复用模块单元
// 以首字母大写来表明可以被包外代码访问
// 2. 代码的package 可以和所在的目录不一致
// 3. 同一个目录里的Go 代码的package 要保持一致
// * 通过go get 来获取远程的依赖： go get -u 强制从网络更新远程依赖
// * 注意代码在Github 上的组织形式， 以适应go get: 直接以代码路径开始， 不要有src。 从源码开始提交，不包括src, 比如从这个项目的src 之下提交到github

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

//小写的方法声明没法再包外被访问到
func getMultiply(a int, b int) int {
	return a * b
}
func GetMultiply(a int, b int) int {
	return a * b
}

// 大写的方法声明可以在包外被访问到
func GetFibonacci(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}
