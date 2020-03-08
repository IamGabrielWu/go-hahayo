package pointer

import (
	"testing"
)

//var var_name *var-type  声明变量的指针  var_name 变量名， 变量类型
var ip *int     /* pointer to an integer */
var fp *float32 /* pointer to a float */
// &a // 访问变量指针地址
// *ip // 访问指针对应的变量的值
func TestPointer(t *testing.T) {
	var a int = 10
	t.Log("内存地址", &a) // pointer 只是内存地址，
}

func TestPointer01(t *testing.T) {
	var ptr *int
	t.Log("The value of ptr is", ptr)
}
