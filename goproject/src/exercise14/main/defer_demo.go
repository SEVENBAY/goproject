package main
import (
	"fmt"
)


func test() {
	fmt.Println("我在test()中")
}


// func init() {
// 	fmt.Println("我在init()中")
// }


func main() {
	var n1 = 1
	var n2 = 2
	defer fmt.Println("n1=", n1)
	defer fmt.Println("n2=", n2)
	defer test()

	n1++
	n2++
	fmt.Println("n1~=", n1)
	fmt.Println("n2~=", n2)
	fmt.Println("main()结束了...")

}