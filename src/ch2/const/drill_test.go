package constant_test

import (
	"fmt"
	"testing"
)

func TestDrill(t *testing.T) {
	const (
		Monday    = iota + 1
		Tuesday   //2
		Wednesday //3
		Thursday  //4
		Friday    //5
		Saturday  //6
		Sunday    //7
	)
	fmt.Println("Monday", Monday)
	fmt.Println("Tuesday", Tuesday)
	fmt.Println("Wednesday", Wednesday)
	fmt.Println("Thursday", Thursday)
	fmt.Println("Friday", Friday)
	fmt.Println("Saturday", Saturday)
	fmt.Println("Sunday", Sunday)

	fmt.Println("##################")
	const (
		Readable   = 1 << iota //1
		Writable               //2
		Executable             //4
	)
	fmt.Println("Readable", Readable)
	fmt.Println("Writable", Writable)
	fmt.Println("Executable", Executable)

}
