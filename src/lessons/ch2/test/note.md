### 编写测试程序
1. 源码文件以 _test 结尾: xxx_test.go
2. 测试方法名以Test 开头: func TestXXX(t *testing.T) {...}
大写的方法代表包外可以访问
### 变量赋值
与其他编程语言的差异
1. 赋值可以自动进行类型推断
2. 在一个赋值语句中可以对多个变量进行同时赋值
### 使用测试方法传参 t *testing.T