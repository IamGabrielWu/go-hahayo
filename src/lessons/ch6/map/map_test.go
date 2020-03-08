package map_test

import "testing"

// must use make to create a map
// Map 声明
// m := map[string]int{"one": 1, "two": 2}
// m1 := map[string]int{}
// 赋值
// m1["one"]=1
// m2 := make(map[string]int, 10 /*initial capacity*/)
// //为什么不初始化len ?
// len的单元格元素都可以赋值为0， 但是Map不行
// map 没法

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

// 对于key 不存在时候， value 默认为0 值， 所以go语言需要判断是否为0

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 2
	t.Log(m1[2])

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
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	//
	for k, v := range m1 {
		t.Log(k, v)
	}
}
func TestMap(t *testing.T) {
	var tech_map = make(map[string]string)
	tech_map["java"] = "4 years"
	tech_map["shell"] = "2 years"
	tech_map["python"] = "4 years"
	for indx, tech := range tech_map {
		t.Log(indx, tech)
	}

	delete(tech_map, "java")
	t.Log("=============")
	for indx, tech := range tech_map {
		t.Log(indx, tech)
	}
}
