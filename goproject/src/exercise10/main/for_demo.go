package main
import (
	"fmt"
)



func main() {
	//输入用户名和密码，判断用户名是否为“张三”，密码是否为1234，如果是，提示登录成功，否则提示登录失败，总共三次机会
	var name string
	var pwd string
	var num int
	// goto FLAG1
	// goto FLAG2
	for i:=3; i>0; i-- {
		fmt.Print("请输入用户名：")
		fmt.Scanln(&name)
		fmt.Print("请输入密码：")
		fmt.Scanln(&pwd)
		if name == "张三" && pwd == "1234" {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Println("登录失败")
			left_chance := i - 1
			if left_chance == 0 {
				fmt.Println("很遗憾，没有机会了")
				break
			}
			fmt.Printf("还有%v次机会\n", left_chance)
		}
	}

	// FLAG1:
	//判断一个整数是否是水仙花数，所谓水仙花数是指一个3位数，其各个位上数字立方和等于其本身。例如: 153 = 1*1*1 + 3*3*3 +5*5*5
	for {
		fmt.Print("请输入一个三位数整数：")
		fmt.Scanln(&num)
		//输入0退出
		if num == 0 {
			break
		}
		//判断范围
		if num < 100 || num >= 1000 {
			fmt.Println("输入有误，请重新输入")
			continue
		}
		//将输入的数字拆分
		num_bai := num / 100
		num_shi := num % 100 / 10
		num_ge := num % 10
		if num == num_bai * num_bai * num_bai + num_shi * num_shi * num_shi + num_ge * num_ge * num_ge {
			fmt.Printf("%v是一个水仙花数\n", num)
		} else {
			fmt.Printf("%v不是一个水仙花数\n", num)
		}
	
	}

	// FLAG2:
	//打印出所有的水仙花数
	fmt.Println("所有的水仙花数为：")
	for i:=100; i <= 999; i++ {
		i_bai := i / 100
		i_shi := i % 100 / 10
		i_ge := i % 10
		if i == i_bai * i_bai * i_bai + i_shi * i_shi * i_shi + i_ge * i_ge * i_ge {
			fmt.Println(i)
		}
	}
}