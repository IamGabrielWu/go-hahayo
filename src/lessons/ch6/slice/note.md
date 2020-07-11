### slice 的声明
var s0 []int // 声明slice 与声明数组的差别， 声明slice没有指定长度
s := []int{}
len(s0) -- slice 长度, len 用于查看slice 和 array 长度
cap(s0) -- slice 容积， cap 用于查看 slice, array 或者map 容积
### 访问元素
// 对切片子数组的改变， 切片后的子数组的改变，也会影响到源数组，因为切片是共享的
Q2 := year[3:6]
s0 = append(s0, 1) -- 增加元素
summer[0] = 0 // 这样会影响其他所有数组。 切片后的子数组的改变，也会影响到源数组，因为切片是共享的