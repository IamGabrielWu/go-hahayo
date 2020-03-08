package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry01(t *testing.T) {
	t.Log(Monday, Tuesday)
}
func TestConstantTry02(t *testing.T) {
	a := 7
	t.Log(Readable)
	t.Log(Writable)
	t.Log(Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
