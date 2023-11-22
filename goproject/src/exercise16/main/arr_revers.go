package main
import (
	"fmt"
)


//将给定的数组逆序
func arrRevers(arr *[4]int) {
	len_arr := len(*arr)
	exchange_num := len_arr / 2
	for i := 0; i < exchange_num; i++ {
		tmp := (*arr)[i]
		(*arr)[i] = (*arr)[len_arr - 1 - i]
		(*arr)[len_arr - 1 - i] = tmp
	}
}


//将给定的切片逆序
func sliceRevers(s []int) {
	len_s := len(s)
	exchange_num := len_s / 2
	for i := 0; i < exchange_num; i++ {
		tmp := s[i]
		s[i] = s[len_s - 1 - i]
		s[len_s - 1 - i] = tmp
	}
}


func main() {
	var arr1 = [...]int{1, 2, 3, 4}
	arrRevers(&arr1)
	fmt.Println(arr1)

	var s1 = []int{1, 2, 3}
	s2 := []int{4, 5, 6, 8, 9, 10}
	sliceRevers(s1)
	sliceRevers(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}