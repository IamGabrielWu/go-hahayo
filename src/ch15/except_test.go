package except_test

import (
	"errors"
	"testing"
)

// Go 的错误机制
// 1. 没有异常机制
// 2. error 类型实现了error 接口
// type error interface {
//   Error() string
// }
// errors.New("n must be in the range [0,10]")
// 3. 可以通过errors.New 来快速创建错误实例
// 最佳实践
// 定义不同的错误变量，以便于判断错误类型
// var LessThanTwoError error = errors.New("n must be greater than 2")
// var GreaterThanHundredError error=errors.New("n must be less than 100")
// func TestGetFibonacci(t *testing.T)  {
//   var list []int
//   list,err :=GetFibonacci(-10)
//   if err == LessThanTwoError {
//     t.Error("Need a larger number")
//   }
//   if err == GreaterThanHundredError {
//     t.Error("need a larger number")
//   }
// }

var LessThanTwoError = errors.New("n should not be less than 2")
var LargerThanHundredError = errors.New("n should not be larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
func TestGetFionacci(t *testing.T) {
	if v, err := GetFibonacci(1); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}

	if v, err := GetFibonacci(1000); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}

	if v, err := GetFibonacci(10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}

	if v, err := GetFibonacci(101); err != nil {
		if err == LessThanTwoError {
			t.Log("<2 error")
		}
		if err == LargerThanHundredError {
			t.Log(">100 error")
		}
	} else {
		t.Log(v)
	}
}

// 及早失败，避免嵌套
