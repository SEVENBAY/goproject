package main
import (
	"fmt"
	"goproject/src/exercise25/utils"
)


func main() {
	var m1 utils.MethodUtils
	m1.Print()
	fmt.Println("==================")
	m1.PrintMethod(9, 90)
	area := m1.GetArea(3.14, 5.6)
	fmt.Println("area=", area)
	m1.Printa()
}