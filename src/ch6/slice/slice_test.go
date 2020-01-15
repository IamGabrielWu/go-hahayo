package slice

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int // 声明slice 与声明数组的差别， 声明slice没有指定长度
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := [...]int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	// cap 代表容量
	// len 代表元素个数
	s2 = append(s2, 2)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s)) // 每一次， 达到cap 上限时候， cap 会已有长度的成倍增长
		// --- PASS: TestSliceGrowing (0.00s)
		// slice_test.go:28: 1 1
		// slice_test.go:28: 2 2
		// slice_test.go:28: 3 4
		// slice_test.go:28: 4 4
		// slice_test.go:28: 5 8
		// slice_test.go:28: 6 8
		// slice_test.go:28: 7 8
		// slice_test.go:28: 8 8
		// slice_test.go:28: 9 16
		// slice_test.go:28: 10 16

	}
}

// 对切片子数组的改变， 切片后的子数组的改变，也会影响到源数组，因为切片是共享的
func TestSliceShareMem(t *testing.T) {
	year :=
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 20, 11, 12}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = 0 // 这样会影响其他所有数组。 切片后的子数组的改变，也会影响到源数组，因为切片是共享的
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	// a := []int{1, 2, 3, 4}
	// b := []int{5, 6, 7, 8}
	// if a == b {
	// t.Log("equal") //   (slice can only be compared to nil)
	//
	// }
}
