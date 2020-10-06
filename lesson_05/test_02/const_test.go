package const_test

import "testing"

// 常量
const (
	// 定义连续常量
	Sunday = iota
	Monday
	Tuesday
)

func TestConst(t *testing.T) {
	t.Log(Sunday, Monday, Tuesday)
}
