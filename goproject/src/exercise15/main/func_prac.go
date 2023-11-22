package main
import (
	"fmt"
	"time"
	"math/rand"
)


/*随机生成一个1--100的整数
有十次机会
如果第一次就猜中，提示 “你真是个天才”
如果第2--3次猜中，提示“你很聪明，赶上我了”
如果第4--9次猜中，提示“一般般”
如果最后一次猜中，提示“可算猜对啦”
一次都没猜对，提示“说你点啥好呢” */
func guessNum() {
	var num int
	var i_num int
	//设置种子
	rand.Seed(time.Now().UnixNano())
	num = rand.Intn(13)
	for i := 1; i < 11; i++ {
		fmt.Print("请输出数字：")
		fmt.Scanln(&i_num)
		if i_num == num {
			if i == 1 {
				fmt.Println("你真是个天才")
			} else if i == 2 || i == 3 {
				fmt.Println("你很聪明，赶上我了")
			} else if i >= 4 && i <= 9 {
				fmt.Println("一般般")
			} else if i == 10 {
				fmt.Println("可算猜对啦")
			}
			return
		} else {
			if i < 10 {
				fmt.Println("猜错了，继续猜")
			}
		}
	}
	fmt.Println("说你点啥好呢？机会用完了，没猜对，答案是", num)
}


/*
编写一个函数，判断是打鱼还是晒网
中国有句俗语叫“三天打鱼两天晒网”。如果从1990年1月1起开始执行“三天打鱼两天晒网”。
如何判断在以后的某一天中是“打鱼”还是“晒网”?
*/
func judgeFish() {
	var year int
	var month int
	var day int
	var start_date string = "1990-1-1"
	var end_date string
	format := "2006-1-2"
	//接收用户输入年、月、日，并判断输入正确与否
	for {
		fmt.Print("请输入年、月、日(使用空格隔开)：")
		fmt.Scanf("%d %d %d", &year, &month, &day)
		if month < 1 || month > 12 {
			fmt.Println("月份输入有误！")
			continue
		}
		if day < 1 || day > 31 {
			fmt.Println("日输入有误！")
			continue
		}
		break
	}
	end_date = fmt.Sprintf("%d-%d-%d", year, month, day)
	start_t, _ := time.Parse(format, start_date)
	end_t, _ := time.Parse(format, end_date)
	duration_t := end_t.Sub(start_t)
	//计算间隔的总天数
	total_d := int(duration_t / (24 * time.Hour)) + 1
	fmt.Printf("经过了%d天\n", total_d)
	tmp := total_d % 5
	if tmp == 1 || tmp == 2 || tmp == 3 {
		fmt.Printf("今天(%d-%d-%d)在打鱼", year, month, day)
	} else if tmp == 4 || tmp == 0 {
		fmt.Printf("今天(%d-%d-%d)在晒网", year, month, day)
	}
}


/*编写一个函数:输出100以内的所有素数(素数就只能被1和本身整除的数)，每行显示5个;并求和*/
func getSu(num int) {
	total := 0
	count := 0
	INDEX:
	for i := 1; i <= num; i++ {
		for j := 2; j < i; j++ {
			value := i % j
			if value == 0 {
				continue INDEX
			}
		}
		total += i
		if count == 5 {
			fmt.Println("")
			count = 0
		}
		count ++
		fmt.Printf("%d\t", i)
	}
	fmt.Println("\n所有素数的和为：", total)
}


func main() {
	// guessNum()
	// judgeFish()
	getSu(1000)
}