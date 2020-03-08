### 循环
与其他语言的差异
* Go 语言仅支持循环关键字 for
for (j:=7;j<=9;j++)
while 条件循环
n :=0
for n<5 {
  n++
  fmt.Println(n)
}
无限循环
while(true)
n:=0
for {

}
### if 条件
if condition {

}else {

}

if condition-1 {

} else if condition-2 {

} else {

}
1. condition 表达式结果必须是布尔值
2. 支持变量赋值
if var declaration; condition {

}

### switch 条件. Go 语言里默认不需要break
switch os:= runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("linux.")
  default:
    fmt.Println("%s.", os)
}
