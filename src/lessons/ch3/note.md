### 基本数据类型
* bool
* string
* int int8 int16 int32 int64
* uint uint8 uint16 uint 32 uint64 uintptr
* byte // alias for uint8
* rune // alias for int32 represents a Unicode code point
* float32 float64
* complex64 complex128
### 类型转换
> 与其他语言的区别 （隐式转换的意思是小的类型可以往大的类型转换，因为不会丢失数据）
1. Go 语言不支持隐式类型转换
2. 别名和原有类型也不能进行隐式类型转换
### 类型的预定义的值
1. math.MaxInt64
2. math.MaxFloat64
3. math.MaxUint32
### 指针类型
1. 不支持指针运算
2. string 是值类型，其默认的初始化值为空字符串，而不是nil
### 显式的类型转换的方法
b = int64(a)
b = string(a)