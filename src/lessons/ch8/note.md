### strings 包
strings.Split(s, ",")
strings.Join(parts, "-")

### strconv 包
包strconv主要实现对字符串和基本数据类型之间的转换。基本数据类型包括：布尔、整型（包括有/无符号、二进制、八进制、十进制和十六进制）和浮点型等。

### 字符串
// 与其他语言区别
// 1. string 是数据类型， 不是引用或者指针
// 2. string 是只读d byte slice, len 函数可以它所包含的byte 数
// 3. string 的byte 数组可以存放任何数据
// 4. Unicode 是一种字符集 (code point)
// 5. UTF8 是unicode 的存储实现（转换为字节序列的规则）

//rune 取出字符串里的unicode， 返回unicode 的slice
// string 是不可变的byte slice