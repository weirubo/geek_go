package func_test

import (
	"fmt"
	"math/rand"
	"testing"
)

// 函数

// 多返回值
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFunc(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

// 可变长参数
func Sum(opt ...int) int {
	ret := 0
	for _, v := range opt {
		ret += v
	}
	return ret
}

func TestFuncValues(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

// defer 延迟函数
func TestDefer(t *testing.T) {
	// defer 延迟函数，在函数中最后执行
	defer func() {
		fmt.Println("Game Over!")
	}()
	fmt.Println("Hello Golang!")
	// 程序 panic，defer 仍正常执行
	// panic("die!")
}
