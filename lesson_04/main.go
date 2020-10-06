package main

import (
	"fmt"
	"os"
)

// 接收命令行参数

func main() {
	// Go 语言 main 函数不能接收参数，不能有返回值。
	// 接收命令行参数。
	argv := os.Args
	str := "hello,"
	if len(argv) > 1 {
		str := fmt.Sprintf(str+"%v", argv[1])
		fmt.Println(str)
	} else {
		// main 函数程序退出。
		os.Exit(1)
	}
}
