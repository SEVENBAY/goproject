package utils
import (
	_"fmt"
)


func GetSum(n1 float32, n2 float32) float32 {
	num := n1 + n2
	return num
}


func GetSub(n1 float32, n2 float32) float32 {
	num := n1 - n2
	return num
}


func GetSumSub(n1 float32, n2 float32) (float32, float32) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}


//可变参数
func NumAdd(n1 int, args... int) int {
	total := n1
	for _, n := range args {
		total += n
	}
	return total
}


//引用参数
func SwapValue(a *float32, b *float32) {
	t := *a
	*a = *b
	*b = t
}

