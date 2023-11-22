package main
import (
	"fmt"
)


func main() {
	fmt.Println("age=", age)
	fmt.Println("name=", name)
}


//初始化函数
func init() {
	fmt.Println("我是初始化函数")
	age = 23
	name = "tom"
}


var (
	age int
	name string
)
