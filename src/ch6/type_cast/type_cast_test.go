package type_cast

import "testing"

func TestTypeCast(t *testing.T) {
	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum) / float32(count)
	t.Log(mean)
}
