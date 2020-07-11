### 通常为了避免对象被拷贝，我们采用指针方式调用
### 所有参数都是有值的传递， slice, map, channel 会传引用的错觉 （看起来传参数是指针， 其实是值的传递）
https://www.geeksforgeeks.org/function-arguments-in-golang/

// 第一种定义方式在实例对应方法被调用时， 实例的成员会进行复制
// func (e Employee) String() string {
// 	fmt.Printf("\nAddress is %x ", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }

// 通常情况下为了避免内存的拷贝，我们使用第二种定义方式
func (e *Employee) String() string {
	fmt.Printf("\nAddress is %x ", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}