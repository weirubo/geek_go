package interface_test

import (
	"fmt"
	"testing"
)

// 接口，空接口和类型断言
type Animal interface {
	Run() string
}

type Dog struct {
	Name string
}

// 实现 Animal 接口的方法
func (d *Dog) Run() string {
	return fmt.Sprintf("%s 正在走路。", d.Name)
}

func TestInterface(t *testing.T) {
	dog1 := Dog{Name: "big yellow"}
	t.Log(dog1.Run())
}

// 空接口和类型断言
func typeOf(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("int 类型", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("string 类型", s)
	}
	fmt.Println("Unknow Type")
}

func TestType(t *testing.T) {
	typeOf(10)
	typeOf("golang")
}
