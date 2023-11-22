package main
import (
	"fmt"
)


//全局匿名函数
var g_func = func (n1 int, n2 int) int {
	return n1 + n2
}


//定义一个函数，以函数作为参数
func g_my_func(f func(int, int) int, n1 int, n2 int) int {
	res := f(n1, n2)
	return res
}


func main() {
	//匿名函数直接调用
	res := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res)

	//匿名函数赋给变量
	my_func := func (n1 int, n2 int) int {
		return n1 * n2
	}
	res1 := my_func(9, 9)
	fmt.Println("res1=", res1)

	//调用全局匿名函数
	res2 := g_func(11, 12)
	fmt.Println("res2=", res2)

	//调用以函数作为参数的函数
	res3 := g_my_func(g_func, 22, 33)
	fmt.Println("res3=", res3)
}