package main
import (
	"fmt"
)


type func_t func(int, int) int


func add_f(n1 int, n2 int) int {
	return n1 + n2
}


func g_f(f func_t, n1 int, n2 int) int {
	return f(n1, n2)
}


func main() {
	var a int = 12
	var b int = 21

	res := g_f(add_f, a, b)
	fmt.Println("res=", res)

}