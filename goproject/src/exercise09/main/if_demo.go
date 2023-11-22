package main
import (
	"fmt"
)


func main() {
	//分支语句使用
	//输入学生分数，根据分数显示不同的信息，0-60：不及格，60-80：良好，80-90：优秀，90-100：顶尖，其他：输入有误
	var score float32
	fmt.Print("请输入分数：")
	fmt.Scanf("%f", &score)
	if score < 60 {
		fmt.Println("不及格")
	} else if score >= 60 && score < 80 {
		fmt.Println("良好")
	} else if score >= 80 && score < 90 {
		fmt.Println("优秀")
	} else if score >= 90 && score <= 100 {
		fmt.Println("顶尖")
	} else {
		fmt.Println("输入有误")
	}

}