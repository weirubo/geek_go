package slice_test

import "testing"

// 切片
func TestSliceInit(t *testing.T) {
	// 声明切片变量
	var s0 []int
	t.Log(len(s0), cap(s0))
	// 追加
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	// 声明切片变量，并初始化
	s1 := []int{1, 2, 3, 4, 5}
	t.Log(len(s1), cap(s1))

	// 内置函数 make，声明切片变量，长度为 3，容量为 5
	// len 个元素被初始化为零值，未初始化的元素不可以访问。
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])

	// 不设置容量，默认容量与长度相等
	s3 := make([]int, 2)
	t.Log(len(s3), cap(s3))
}

func TestSliceGrowing(t *testing.T) {
	s4 := []int{}
	// 切片空间不够用时，随元素长度成倍增长。
	for i := 0; i < 10; i++ {
		s4 = append(s4, i)
		t.Log(len(s4), cap(s4))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}
