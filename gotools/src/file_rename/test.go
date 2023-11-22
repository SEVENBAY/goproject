package main
import (
	"fmt"
	"strings"
	_"bufio"
	_"os"
)


func main() {
	s := "我是文件名.txt"
	r := strings.Replace(s, "我是", "", -1)
	fmt.Println(r)

	var name string
	fmt.Printf("请输入：")
	fmt.Scanln(&name)
	fmt.Printf("%v, %T\n", name, name)
	fmt.Println("========", len(name))
	r2 := strings.Replace(s, "我是", name, -1)
	fmt.Println("==============", r2)


	// var gender string
	// fmt.Printf("输入吧：")
	// sc := bufio.NewReader(os.Stdin)
	// res, _, _ := sc.ReadLine()
	// fmt.Printf("%v, %T\n", string(res), string(res))
	// fmt.Println("========", len(res))
}