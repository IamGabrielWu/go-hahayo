package local

import (
	"ch17/series" // 包的导入路径
	"testing"
)

func TestLocalPackage(t *testing.T) {
	t.Log(series.GetFibonacci(100))
	// t.Log(series.getMultiply(10, 20)) //小写的方法声明没法在保外被访问到
	t.Log(series.GetMultiply(10, 20))
}

// # init 方法
// * 在main 被执行前， 所有依赖的package 的init 方法都会被执行
// * 不同包的init 函数按照包导入的依赖关系决定执行顺序
// * 每个包可以有多个init 函数
// * 包的每个源文件也可以有多个init 函数， 这点比较特殊
