package array_test

import "testing"

// 数组

func TestArrayInit(t *testing.T) {
	// 声明数组变量
	var arr [3]int
	// 声明数组变量并初始化
	arr1 := [4]int{1, 2, 3, 4}
	// 声明数组变量并初始化，根据初始化值自动定义数组长度
	arr2 := [...]int{1, 2, 3, 4, 5}
	// 根据数组下标，设置数组元素
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 5, 7, 9}
	// 遍历数组
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	// for range
	for key, val := range arr3 {
		t.Log(key, val)
	}
}

func TestArraySection(t *testing.T) {
	arr4 := [...]int{1, 2, 3, 4, 5, 6}
	// 截取数组前 3 个元素
	arrSec1 := arr4[:3]
	t.Log(arrSec1)
	// 截取数组下标为 3 的元素后的所有元素
	arrSec2 := arr4[3:]
	t.Log(arrSec2)
	// 截取数组的所有元素
	arrSec3 := arr4[:]
	t.Log(arrSec3)
	// 截取数组下标 1-3 的元素
	// 包含开始元素，不包含结束元素
	arrSec4 := arr4[1:4]
	t.Log(arrSec4)
}
