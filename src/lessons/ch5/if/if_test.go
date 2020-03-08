package condition

import "testing"

func TestIfMultiSet(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
}

// go 语言支持多返回值
// func TestIf(t *testing.T) {
// 	if v, err := someFun(); err == nil {
// 		t.Log("1==1")
// 	} else {
// 		t.Log()
// 	}
// }
// func someFun() {
//
// }

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 4:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 4:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")

		}
	}
}
