### 退出返回值
与其他语言的差别
1. Go 中的main 函数不支持返回值
2. 通过os.Exit 来返回状态
### 获取命令行参数
与其他语言的差异
1. main 函数不支持传入参数
> func main(arg [] string)
2. 在程序中直接通过os.Args获取命令行参数 