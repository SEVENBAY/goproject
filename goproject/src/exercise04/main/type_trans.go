package main
import (
	"fmt"
	"strconv"
)


func main() {
	// 其他类型转字符串
	var i int = 100
	var j float32 = 23.45
	var k bool = true
	//其他类型转字符串
	//第一种方式
	fmt.Println("=========第一种方式===================")
	i_str1 := fmt.Sprintf("%d", i)
	j_str1 := fmt.Sprintf("%f", j)
	k_str1 := fmt.Sprintf("%t", k)
	fmt.Printf("i转成字符串是%v,类型是%T\n", i_str1, i_str1)
	fmt.Printf("j转成字符串是%v,类型是%T\n", j_str1, j_str1)
	fmt.Printf("k转成字符串是%v,类型是%T\n", k_str1, k_str1)
	//第二种方式
	fmt.Println("==========第二种方式=================")
	i_str2 := strconv.FormatInt(int64(i), 10)
	j_str2 := strconv.FormatFloat(float64(j), 'f', 5, 32)
	k_str2 := strconv.FormatBool(k)
	fmt.Printf("i转成字符串是%v,类型是%T\n", i_str2, i_str2)
	fmt.Printf("j转成字符串是%v,类型是%T\n", j_str2, j_str2)
	fmt.Printf("k转成字符串是%v,类型是%T\n", k_str2, k_str2)
	

	// 字符串转其他类型
	fmt.Println("=========字符串转其他类型===============")
	i_int, _ := strconv.ParseInt(i_str2, 10, 32)
	j_float, _ := strconv.ParseFloat(j_str2, 32)
	k_bool, _ := strconv.ParseBool(k_str2)
	fmt.Printf("i_int=%v,i_int的类型是%T\n", i_int, i_int)
	fmt.Printf("j_float=%v,j_float的类型是%T\n", j_float, j_float)
	fmt.Printf("k_bool=%v,k_bool的类型是%T\n", k_bool, k_bool)

}