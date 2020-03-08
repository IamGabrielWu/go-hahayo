package fib

import (
	"fmt"
	"testing"
)

func TestDrill(t *testing.T) {
	letters := []string{"a", "b", "c", "d", "e", "f", "g"}
	exchangeLettersInArray(&letters)
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	exchangeLettersInArray(&numbers)

	var arr = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	var pf *[6]int = &arr
	for _, i := range pf {
		fmt.Println(i)
	}
	//指针数组
	// pfArr := [...]*int{&x, &y}
	fmt.Println(pf)
	x, y := 1, 2
	pfArr := []*int{&x, &y}
	fmt.Println(pfArr)
}

// 倒序所有字符
// 数组指针传递， &letters，--- 这是指针， 方法名 letters *[]string， 指定传入方式为指针
// 访问指针表示的数组的内部值 (*letters)[i] --- (*letters) 为指针代指向的值
func exchangeLettersInArray(letters *[]string) {
	mid := len(*letters) / 2
	last := len(*letters) % 2
	fmt.Println(mid, last)
	if last == 0 {
		for i := 0; i < len(*letters); i++ {
			if i == (mid - 1) {
				break
			}
			(*letters)[i], (*letters)[len((*letters))-1-i] = (*letters)[len((*letters))-1-i], (*letters)[i]
			fmt.Println((*letters))
		}
	} else {
		for i := 0; i < len((*letters)); i++ {
			if i == mid {
				break
			}
			(*letters)[i], (*letters)[len((*letters))-1-i] = (*letters)[len((*letters))-1-i], (*letters)[i]
			fmt.Println((*letters))
		}
	}
	fmt.Println((*letters))
}
