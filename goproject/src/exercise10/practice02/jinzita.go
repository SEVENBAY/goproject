package main
import (
	"fmt"
)


func main() {
	//打印金字塔
	var level int = 40  //金字塔层数
	for i:=1; i<=level; i++ {
		for k:=level-i; k>0; k-- {
			fmt.Print(" ")
		}
		for j:=2*i-1; j>0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	//打印空心金字塔
	for i:=1; i<=level; i++ {
		for k:=level-i; k>0; k-- {
			fmt.Print(" ")
		}
		for j:=2*i-1; j>0; j-- {
			if i != level {
				if j == 2 * i -1 || j == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println("")
	}

}