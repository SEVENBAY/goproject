package main
import (
	"fmt"
)


func demo1() {
	//使用 switch 把小写类型的 char型转为大写(键盘输入)。只转换a,b,c,d,e.其它的输出“other”。
	var i byte
	fmt.Print("请输入字母：")
	fmt.Scanf("%c", &i)
	i_new := i - 32
	switch i {
	case 'a':
		fmt.Printf("%c\n", i_new)
	case 'b':
		fmt.Printf("%c\n", i_new)
	case 'c':
		fmt.Printf("%c\n", i_new)
	case 'd':
		fmt.Printf("%c\n", i_new)
	case 'e':
		fmt.Printf("%c\n", i_new)
	default:
		fmt.Println("other")
	}	

}


func demo2() {
	//对学生成绩大于60分的，输出“合格”。低于60分的，输出“不合格”。(注:输入的成绩不能大于100)
	var i float32
	fmt.Print("请输入学生成绩：")
	fmt.Scanln(&i)

	switch {
	case i > 100:
		fmt.Println("输入有误")
	case i >= 60:
		fmt.Println("合格")
	case i < 60:
		fmt.Println("不合格")
	}
}


func demo3() {
	//根据用户指定月份，打印该月份所属的季节。3,4,5 春季 6,7,8 夏季 9,10,11 秋季12,1,2 冬季
	var i byte
	fmt.Print("请输入月份：")
	fmt.Scanln(&i)

	switch i {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12, 1, 2:
		fmt.Println("冬季")
	default:
		fmt.Println("输入有误")
	}
}


func demo4() {
	//根据用户输入显示对应的星期时间(string)，如果“星期一”，显示“干煽豆角”如
	//果“星期二”，显示“醋溜土豆”如果“星期兰”，显示“红烧狮子头”如果“星
	//期四”显示“油炸花生米” 如果“星期五”，显示“蒜蓉扇贝” 如果“星期六”，
	//显示“东北乱炖”，如果“星期日”，显示“大盘鸡”
	var i string
	fmt.Print("请输入今天是星期几：")
	fmt.Scanln(&i)
	switch i {
	case "星期一":
		fmt.Println("干煸豆角")
	case "星期二":
		fmt.Println("醋溜土豆")
	case "星期三":
		fmt.Println("红烧狮子头")
	case "星期四":
		fmt.Println("油炸花生米")
	case "星期五":
		fmt.Println("蒜蓉扇贝")
	case "星期六":
		fmt.Println("东北乱炖")
	case "星期日":
		fmt.Println("大盘鸡")
	default:
		fmt.Println("西北风")
	}


}



func main() {
	demo1()
	demo2()
	demo3()
	demo4()
}