package main
import (
	"fmt"
)


func main() {
	//打印99乘法表
	for i:=1; i<=9; i++ {
		for j:=1; j<=9; j++ {
			fmt.Printf("%v x %v = %v\t", j, i, (i * j))
			if i == j {
				fmt.Println("")
				break
			}
		}
	}



}