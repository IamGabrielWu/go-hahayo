package object_test

import (
	"fmt"
	"testing"
	"unsafe"
)

// Is Go an object-oriented language ?
// Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concept of "interface" in Go provides a different approach that we believe is easy to use and in some ways more general.
// Also, the lack of a type hierarchy makes "objects" in Go feel much more lightweight than in lanuage such as C++ or Java
type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 20}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"

	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)   // struct 类型
	t.Logf("e2 is %T", e2) // 返回指针类型
	t.Logf("e is %T", &e)  // 加上& 返回指针类型
}

// 行为(方法)定义
// 与其他语言的区别
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
func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("\nTestStructOperations is %x", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
