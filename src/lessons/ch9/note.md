// 函数是一等公民
// 与其他语言的区别
// 1.可以有多个返回值
// 2. 所有参数都是有值的传递， slice, map, channel 会传引用的错觉
// 3. 函数可以作为变量的值
// 4. 函数可以作为参数和返回值

a, _ = returnMultiVal() // _ 下划线， 忽略其中一个返回值
https://www.geeksforgeeks.org/function-arguments-in-golang/
所有参数都是有值的传递, 区别在于是复制一份原值传递给func 还是把已有的值传递给func 处理
