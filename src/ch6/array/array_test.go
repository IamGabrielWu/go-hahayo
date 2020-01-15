package array

import "testing"

func TestArrayInt(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5} // 自动按照元素个数初始化数组长度
	t.Log(arr[1], arr1[2], arr2[3])
}
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 5, 7}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 3, 5, 7}
	arr3_sec := arr3[:3]
	t.Log(arr3_sec)
	arr_sec := arr3[3:]
	t.Log(arr_sec)
}
