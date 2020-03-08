package string_test

import (
	"testing"
)

// 字符串
// 与其他语言区别
// 1. string 是数据类型， 不是引用或者指针
// 2. string 是只读d byte slice, len 函数可以它所包含的byte 数
// 3. string 的byte 数组可以存放任何数据
// 4. Unicode 是一种字符集 (code point)
// 5. UTF8 是unicode 的存储实现（转换为字节序列的规则）

//rune 取出字符串里的unicode， 返回unicode 的slice
func TestString(t *testing.T) {
	var s string
	t.Log(s) // 默认初始化控制符
	s = "hello"
	t.Log(len(s))
	// s[1] = '3'         // string 是不可变的byte slice
	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	// s = "\xE4\xBB\xA5\xFF"
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s))
	c := []rune(s)
	// t.Log("rune sie:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c)
	t.Logf(" 中 UTF8 %x", s)
}

// 常用的字符串函数
// 1. strings 包
// 2. strconv 包

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		// c 表示字符串, d 编码, x 十六进制
		t.Logf("%[1]c %[1]d %[1]x", c)
	}
}
