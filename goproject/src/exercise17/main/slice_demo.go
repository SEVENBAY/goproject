package main
import (
	"fmt"
)


func main() {
	var s1 []int = make([]int, 4, 8)
	var s2 []int = []int{1, 2, 3, 10, 12}
	s3 := []string{"张三", "李四", "王五"}
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("s3=", s3)
	fmt.Println()
	copy(s1, s2)
	fmt.Println("s1=", s1)
	s4 := append(s3, "赵六", "高七", "王八")
	fmt.Println("s4=", s4, "s3=", s3)
	s5 := append(s1, s2...)
	fmt.Println("s5=", s5)
	fmt.Println("========================")
	//遍历切片s4
	for i, v := range s4 {
		fmt.Printf("第%d个元素为：%v\n", i, v)
	}
	fmt.Println("=======================")
	for i := 0; i < len(s4); i++ {
		fmt.Printf("第%d个元素为：%v\n", i, s4[i])
	}

	var arr [6]string = [...]string{"张三", "李四", "王五", "赵六", "高七", "王八"}
	s6 := arr[:5]
	fmt.Printf("s6=%v，s6的类型=%T\n", s6, s6)
	arr[0] = "tom"
	fmt.Println("arr=", arr, "s6=", s6)
	copy(s6, []string{"aa", "bb"})
	fmt.Println("arr=", arr, "s6=", s6)
	fmt.Printf("arr的地址=%p，s6的地址=%p", &arr, &s6[0])
}