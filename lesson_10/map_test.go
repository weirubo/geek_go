package map_test

import "testing"

// map

func TestMapInit(t *testing.T) {
	// 声明 map 变量
	m1 := map[int]int{}
	t.Log(m1, len(m1))
	m1[5] = 5
	t.Log(m1, len(m1))

	m2 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m2, m2[2], len(m2))

	// 设置 map 的容量
	m3 := make(map[int]int, 5)
	t.Log(m3, len(m3))
}

func TestMapNotExist(t *testing.T) {
	m4 := map[int]int{}
	t.Log(m4[1])
	m4[2] = 0
	t.Log(m4[2])

	m4[3] = 0
	// 判断 key 是否是零值
	if v, ok := m4[3]; ok { // 存在
		t.Log(v)
	} else {
		t.Log("Not exits")
	}
}

func TestMapTravel(t *testing.T) {
	m5 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	// 遍历 map
	for key, val := range m5 {
		t.Log(key, val)
	}
}
