package struct_test

import (
	"fmt"
	"testing"
	"unsafe"
)

// 结构体和方法

type User struct {
	ID   int
	Name string
	Age  int
}

func TestStruct(t *testing.T) {
	// 声明结构体变量
	user1 := User{1, "lucy", 16}
	t.Log(user1)
	user2 := User{ID: 2, Name: "lily", Age: 17}
	t.Log(user2)
	// 返回指针，相当于 user3 := &User{}
	user3 := new(User)
	user3.ID = 3
	user3.Name = "frank"
	user3.Age = 18
	t.Log(user3)
	t.Logf("user1 type:%T", user1)
	t.Logf("user3 type:%T", user3)
}

func (u User) String() string {
	fmt.Printf("%x\n", unsafe.Pointer(&u.Name))
	return fmt.Sprintf("ID:%d-Name:%s-Age:%d", u.ID, u.Name, u.Age)
}

/* func (u *User) String() string {
	fmt.Printf("%x\n", unsafe.Pointer(&u.Name))
	return fmt.Sprintf("ID:%d/Name:%s/Age:%d", u.ID, u.Name, u.Age)
} */

func TestMethod(t *testing.T) {
	/* user1 := User{1, "lucy", 16}
	t.Log(user1.String()) */
	user1 := &User{1, "lucy", 16}
	t.Log(user1.String())
	fmt.Printf("%x\n", unsafe.Pointer(&user1.Name))
}
