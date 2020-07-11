### map 定义
var xxx = map[int]bool{}, 即 map[type]type{}
m2 := make(map[string]int, 容量capacity)
var tech_map = make(map[string]string)
### delete value in map
delete(my_map, 1)
### set value in map
my_map[1] = true
### iterator map
for k, v := range m1 {
    t.Log(k, v)
}
### map 对于key 不存在时候， value 默认为0 值， 所以go语言需要判断是否为0 (很重要)
```
	// 判断key 是否存在
	m1[3] = 10
	if v, ok := m1[3]; ok {
		t.Logf("Key 3 exist and value is %d", v)
	} else {
		t.Logf("Key 3 not exist %d", v)
	}

	// Map 元素的访问
	// 与其他语言的差异
	// 在访问的Key 不存在的时， 仍会返回零值， 不能通过返回nil来判断元素是否存在
```