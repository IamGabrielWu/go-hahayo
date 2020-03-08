package operator_test

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	fmt.Println(a == b)
	fmt.Println(a == d)
	// t.Log(a == b)
	// t.Log(a == c) // 数组内元素个数不同，无法比较
	// t.Log(a == d)
}
