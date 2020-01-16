package map_test

import "testing"

// must use make to create a map

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
