package main
import (
	"fmt"
	"time"
	"math/rand"
	"math"
	"strconv"
)


//随机生成10个整数(1 100的范围)保存到数组，并倒序打印以及求平均值、求最大值和最小值的下标、并查找里面是否有55
func demo1() {
	var arr [10]int
	var max_n_index int = 0
	var min_n_index int = 0
	var sum int = 0
	var b_55 bool = false

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(101)
	}
	fmt.Println("原arr=", arr)

	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		sum += num
		if num == 55 {
			b_55 = true
		}
		if arr[max_n_index] < num {
			max_n_index = i
		}
		if arr[min_n_index] > num {
			min_n_index = i
		}
		fmt.Printf("%v ", num)
	} 
	fmt.Println()
	fmt.Println("最大值的下标为：", max_n_index)
	fmt.Println("最小值的下标为：", min_n_index)
	fmt.Println("平均值为：", sum / len(arr))
	if b_55 == true {
		fmt.Println("arr中存在数值55")
	} else {
		fmt.Println("arr中不存在数值55")
	}

}


//已知有个排序好(升序)的数组，要求插入一个元素，最后打印该数组，顺序依然是升序
func demo2(a [5]int, e int) {
	arr := a[:]
	var new_s = make([]int, 6)
	if e <= arr[0] {
		new_s[0] = e
		for i := 0; i < len(arr); i++ {
			new_s[i + 1] = arr[i]
		}
	} else if e >= arr[4] {
		copy(new_s, arr)
		new_s[len(arr)] = e
	} else {
		new_s_index := 0
		for i:= 0; i < len(arr)-1; i++ {
			new_s[new_s_index] = arr[i]
			if new_s_index == i && arr[i] <= e && e <= arr[i+1] {
				new_s_index += 1
				new_s[new_s_index] = e
			}
			new_s_index ++
		}
		new_s[new_s_index] = arr[len(arr)-1]
	}
	fmt.Println("插入元素后的数组=", new_s)
}


//定义一个3行4列的二维数组，逐个从键盘输入值，编写程序将四周的数据清0
func demo3() {
	var arr [3][4]int
	for i := 0; i < len(arr); i++ {
		fmt.Println("请分别输入第", i + 1, "个数组的值")
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("请输入第%v个值：", j + 1)
			fmt.Scanln(&arr[i][j])
		}
	}
	//输出
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Printf("%v\t", v2)
		}
		fmt.Println()
	}
	//四周数据清零
	for i, v1 := range arr {
		//将第一行和最后一行清零
		if i == 0 || i == len(arr) - 1 {
			for j := 0; j < len(v1); j++ {
				arr[i][j] = 0
			}
		} else {
			//中间行将首尾的数字清零
			arr[i][0] = 0
			arr[i][len(v1) - 1] = 0
		}

	}
	//输出
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Printf("%v\t", v2)
		}
		fmt.Println()
	}
}


//定义一个4行4列的二维数组逐个从键盘输入值，然后将第1行和第4行的数据进行交换，将第2行和第3行的数据进行交换
func demo4() {
	var arr [4][4]int
	for i := 0; i < len(arr); i++ {
		fmt.Println("请分别输入第", i + 1, "个数组的值")
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("请输入第%v个值：", j + 1)
			fmt.Scanln(&arr[i][j])
		}
	}
	//输出
	fmt.Println("交换前：")
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Printf("%v\t", v2)
		}
		fmt.Println()
	}
	//交换
	temp := arr[0]
	arr[0] = arr[3]
	arr[3] = temp
	temp = arr[1]
	arr[1] = arr[2]
	arr[2] = temp

	fmt.Println("交换后：")
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Printf("%v\t", v2)
		}
		fmt.Println()
	}

}


//定义一个数组，并给出8个整数，求该数组中大于平均值的数的个数，和小于平均值的数的个数。
func demo5() {
	var arr [8]int
	var sum int = 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		n := rand.Intn(100)
		sum += n
		arr[i] = n
	}
	fmt.Println("arr=", arr)

	avr_num := sum / len(arr)
	var avr_low = 0
	var avr_high = 0
	for _, v := range arr {
		if v > avr_num {
			avr_high ++
		}
		if v < avr_num {
			avr_low ++
		}
	}
	fmt.Printf("平均值为 %v，大于平均值的个数为 %v，小于平均值的个数为 %v", avr_num, avr_high, avr_low)
}


/*
跳水比赛，8个评委打分。运动员的成绩是8个成绩取掉一个最高分，去掉一个最低分，剩下的6个分数的平均分就是最后得分。使用一维数组实现如下功能:
(1) 请把打最高分的评委和最低分的评委找出来。
(2) 找出最佳评委和最差评委。最佳评委就是打分和最后得分最接近的评委。最差评委就是打分和最后得分相差最大的。
*/
func demo6() {
	var score [8]float32
	var max_index = 0
	var min_index = 0
	var total float32

	rand.Seed(time.Now().UnixNano())
	//打分
	for i := 0; i < len(score); i++ {
		s := float32(rand.Intn(100)) + rand.Float32()
		s1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", s), 32)
		score[i] = float32(s1)
	}
	fmt.Println("打分结果：", score)
	for i, v := range score {
		total += v
		if score[max_index] < v {
			max_index = i
		}
		if score[min_index] > v {
			min_index = i
		}
	}
	fmt.Println("打最低分的评委编号为：", min_index)
	fmt.Println("打最高分的评委编号为：", max_index)

	//计算平均分
	avr_score := (total - score[min_index] - score[max_index]) / float32(len(score) - 2)
	var best_index int
	var worest_index int
	var best_num = 0.0
	var worest_num = 100.0
	for i, v := range score {
		if i == min_index || i == max_index {
			continue
		}
		//将数组各个元素变成与平均值的差值
		n := math.Abs(float64(v - avr_score))
		if n < worest_num {
			worest_num = n
			best_index = i
		}
		if n > best_num {
			best_num = n
			worest_index = i
		}

	}
	fmt.Printf("评委最终的平均分为：%v，最佳评委编号为：%v，最差评委为：%v", avr_score, best_index, worest_index)
}



func main() {
	demo1()
	fmt.Println()
	var arr = [5]int{10, 20, 30, 40, 50}
	demo2(arr, 29)
	fmt.Println()
	demo3()
	fmt.Println()
	demo4()
	fmt.Println()
	demo5()
	fmt.Println()
	demo6()
}