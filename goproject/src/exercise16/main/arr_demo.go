package main
import (
	"fmt"
)


func main() {
	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7}
	var arr3 = [...]int{9, 10, 11}
	arr4 := [...]int{1:12, 0:13, 2:14}
	fmt.Println("arr1=", arr1)
	fmt.Println("arr2=", arr2)
	fmt.Println("arr3=", arr3)
	fmt.Println("arr4=", arr4)
	//遍历数组arr2
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d]=%d\n", i, arr2[i])
	}
	fmt.Println("")
	//遍历数组arr3
	for i, v := range arr3 {
		fmt.Printf("arr3[%d]=%d\n", i, v)
	}
}