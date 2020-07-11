### 访问指针 以及表示指针
// &xxx -> 访问指针
// function abc(xxx *type){}  *type 表示以指针传入
// *xxx  表示指针所指向的值。 实际的值。
### 指针的声明和使用
//var var_name *var-type  声明变量的指针  var_name 变量名， *变量类型
var ip *int     /* pointer to an integer */
var fp *float32 /* pointer to a float */
// &a // 访问变量指针地址
// *ip // 访问指针对应的变量的值