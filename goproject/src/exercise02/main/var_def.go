package main

import (
	"fmt"
	"unsafe"
)

var (
	a int = 1
	b int = 2
	c     = 3
)

// 入口函数
func main() {
	// 变量定义
	var i int = 2
	var j = 3
	k := 4
	fmt.Println("i=", i, "j=", j, "k=", k)
	// 多变量定义
	fmt.Println("a=", a, "b=", b, "c=", c)
	var a, b, c int
	fmt.Println("a=", a, "b=", b, "c=", c)
	// 查看类型及所占字节大小
	var m = 10
	fmt.Printf("m的类型为%T\n", m)
	fmt.Printf("m所占字节大小为%d\n", unsafe.Sizeof(m))
	// n := "1"
	var n byte = '1'
	fmt.Printf("n的类型为%T\n", n)
	fmt.Printf("n所占字节大小为%d\n", unsafe.Sizeof(n))
	fmt.Printf("n的unicode值为%v", n)
}
