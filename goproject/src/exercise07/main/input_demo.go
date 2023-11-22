package main
import (
	"fmt"
)


func main() {
	//接收标准输入
	var age byte
	var name string
	var sal float32

	// Scanln
	// fmt.Println("请输入姓名：")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)
	// fmt.Println("请输入工资：")
	// fmt.Scanln(&sal)

	//Scanf
	fmt.Println("请输入姓名、年龄、工资，使用空格隔开：")
	fmt.Scanf("%s %d %f", &name, &age, &sal)


	fmt.Printf("此人的信息：name-%v, age-%v, sal-%v", name, age, sal)
}