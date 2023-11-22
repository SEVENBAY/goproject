package main
import (
	"fmt"
)


func main() {
	var a int8
	var b float32
	var c bool
	var d byte
	var pt * int8
	a = 45
	b = 3.1415926
	c = true
	d = 'a'
	pt = &a
	fmt.Printf("a=%v, b=%v, c=%v, d=%v, pt=%v\n", a, b, c, d, *pt)
	fmt.Printf("d的原始值为%c\n", d)
	// 通过指针修改a的值
	*pt = 2
	fmt.Printf("a=%v, pt=%v\n", a, *pt)
	// 类型转换
	var e int
	e = int(b)
	fmt.Printf("e=%v\n", e)
	// 算数运算
	var i float32
	i = 2 * 3.1415926
	var j int
	j = 2 * 3
	fmt.Printf("i=%v, j=%v", i, j)
}