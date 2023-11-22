package utils
import (
	"fmt"
)


type MethodUtils struct{

}


func (this *MethodUtils) Print() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (this *MethodUtils) PrintMethod(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (this *MethodUtils) GetArea(length float64, width float64) float64 {
	return length * width
}


func test() {
	fmt.Println("我在exec1.go中")
}