package var_test

import "testing"

// 变量
func TestVar(t *testing.T) {
	a := 1
	b := 1
	t.Log(a)
	// 打印斐波拉契数列
	for i := 0; i <= 5; i++ {
		t.Log(b)
		/* tmp := a
		a = b
		b = tmp + a */
		// Golang 支持同时声明多个变量，此处可以不使用临时变量。
		a, b = b, a+b
	}
}
