package string_test

import (
	"strconv"
	"strings"
	"testing"
)

// 字符串
func TestString(t *testing.T) {
	// 声明字符串
	var s string
	t.Log(s) // 字符串的零值是空字符串

	s = "golang"
	t.Log(len(s))
	// 字符串是不可变的字节切片
	// s[1] = "g"
	s = "编程"
	t.Log(len(s)) // 字符串的长度是字节长度
	c := []rune(s)
	t.Log(len(c))
	t.Logf("unicode:%x", c[0])
	t.Logf("UTF8:%x", s)
}

func TestStringToRune(t *testing.T) {
	str := "技术社区"
	for _, c := range str {
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFunc(t *testing.T) {
	// 字符串切割为字符切片与字符切片拼接为字符串
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T) {
	// 字符串与整型之间的类型转换
	t.Log("str" + strconv.Itoa(10))
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
