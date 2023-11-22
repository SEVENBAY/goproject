package main
import (
	"fmt"
	"unsafe"
)


func main() {
	var s string = "abcd我要学go"
	fmt.Printf("字符串s的长度为%d(Sizeof)\n", unsafe.Sizeof(s))
	fmt.Printf("字符串s的长度为%d(len)\n", len(s))
	//字符串迭代方式1
	for id, ele := range(s) {
		fmt.Printf("id-%v, ele-%c\n", id, ele)
	}
	fmt.Println("")
	//字符串迭代方式2
	s_new := []rune(s)
	for i:=0;i<len(s_new);i++ {
		fmt.Printf("id-%d, ele-%c\n", i, s_new[i])
	}
}