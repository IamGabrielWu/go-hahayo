package main // 可以不使用目录名叫做main 的目录，但是声明的package 必须是main

import (
	"fmt" //引入代码依赖
	"os"
)

//功能实现
func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello World ", os.Args[1])
	}
	fmt.Println("Hello World")
	os.Exit(126)
}
