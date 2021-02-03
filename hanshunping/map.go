package main

import (
	"fmt"
	"sort"
	"unsafe"
)

/*
	go 映射
*/
func main() {

	testMap()
	testMapSlice()
	testMapSort()
}

func testMap() {
	var m1 = make(map[string]string, 10)
	m1["k1"] = "v1"

	m2 := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	fmt.Println(m2)
	// 查找
	v1 := m2["k1"]
	fmt.Println(v1)

	// map<string, object>
	m3 := make(map[string]interface{}, 10)
	m3["k1"] = "v1"
	m3["k2"] = 2
	fmt.Println(m3)
	delete(m3, "k1")

	// 清空map
	for k, v := range m3 {
		fmt.Print("k =", k, ", v =", v)
		delete(m3, k)
	}
	fmt.Println()

	// 删除m2所有k，使之前引用的map成为垃圾
	m2 = make(map[string]string)
}

/*
	map 切片
*/
func testMapSlice() {
	sm := make([]map[string]interface{}, 4)
	sm[0] = make(map[string]interface{}, 4)
	sm[0]["k1"] = "s1v1"
	sm[0]["k2"] = "s1v2"

	sm[1] = make(map[string]interface{}, 4)
	sm[1]["k1"] = "s2v1"
	sm[1]["k2"] = "s2v2"

	sm10 := make(map[string]interface{}, 4)
	sm10["k1"] = "sm10v1"
	sm = append(sm, sm10)

	fmt.Println(unsafe.Sizeof(sm))
	fmt.Println(sm)
}

/*
	map排序
		1. 遍历map将k放入slice
		2. 排序slice
		3. 遍历slice，获取k对应v
*/
func testMapSort() {
	m2 := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k4": "v4",
		"k3": "v3",
	}
	var ks []string
	for k := range m2 {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	for _, k := range ks {
		fmt.Print(m2[k])
	}
}
