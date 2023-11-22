package main
import (
	"fmt"
	util "exercise11/utils"
)


func main() {
	var n1 float32 = 3.2
	var n2 float32 = 4.3

	sum := util.GetSum(n1, n2)
	fmt.Println("n1 + n2 =", sum)

	sub := util.GetSub(n1, n2)
	fmt.Printf("n1 - n2 =%.2f\n", sub)

	sum1, sub1 := util.GetSumSub(n1, n2)
	fmt.Printf("n1 + n2 =%v\nn1 - n2 =%.2f\n", sum1, sub1)

	//可变参数
	res := util.NumAdd(1, 2, 3, 4, 5)
	fmt.Println("res=", res)

	//引用传递
	util.SwapValue(&n1, &n2)
	fmt.Printf("n1=%v, n2=%v\n", n1, n2)

}