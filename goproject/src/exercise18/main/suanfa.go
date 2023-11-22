package main
import (
	"fmt"
	"time"
	"math/rand"
)


//冒泡排序
func bubbleSort(slice []int) {
	slice_len := len(slice)
	for i := 0; i < slice_len - 1; i++ {
		for j := 0; j < slice_len - 1 - i; j++ {
			if slice[j] > slice[j+1] {
				tmp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = tmp
			}
		}
	}
}


//二分查找
func halfFind(slice []int, value int, start int, end int) {
	if start > end {
		fmt.Printf("元素%v未找到\n", value)
		return
	}
	half := (start + end) / 2
	if value == slice[half] {
		fmt.Printf("元素%v找到了，元素下标=%d", value, half)
		return
	} else if value > slice[half] {
		halfFind(slice, value, half + 1, end)
	} else {
		halfFind(slice, value, start, half - 1)
	}
}



func main() {
	//定义一个切片，随机填充数字，然后进行排序（升序）
	var slice = make([]int, 20)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(slice); i++ {
		num := rand.Intn(100)
		slice[i] = num
	}
	fmt.Println("排序前：slice=", slice)
	bubbleSort(slice)
	fmt.Println("排序后：slice=", slice)

	find_n := rand.Intn(100)
	halfFind(slice, find_n, 0, len(slice)-1)

}