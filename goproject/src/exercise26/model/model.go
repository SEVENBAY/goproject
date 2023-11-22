package model
import (
	"fmt"
)


type Account struct{
	Number string
	Extra float64
	Password string
}

func (this *Account) SetNumber(number string) bool {
	if len(number) < 6 || len(number) > 10 {
		fmt.Println("输入的number不合规，赋值失败")
		return false
	}
	this.Number = number
	return true
}

func (this *Account) SetExtra(extra float64) bool {
	if extra <= 20.0 {
		fmt.Println("输入的extra不合规，赋值失败")
		return false
	}
	this.Extra = extra
	return true
}

func (this *Account) SetPassword(password string) bool {
	if len(password) != 6 {
		fmt.Println("输入的password不合规，赋值失败")
		return false
	}
	this.Password = password
	return true
}

func (this *Account) ShowDetail() {
	fmt.Println("账号信息如下：")
	fmt.Println("Number=", this.Number)
	fmt.Println("Extra=", this.Extra)
	fmt.Println("Password=", this.Password)
}


func NewAccount(number string, extra float64, password string) *Account {
	//分别判断各个字段是否合规
	if len(number) < 6 || len(number) > 10 {
		fmt.Println("输入的number不合规")
		return nil
	}
	if extra <= 20.0 {
		fmt.Println("输入的extra不合规")
		return nil
	}
	if len(password) != 6 {
		fmt.Println("输入的password不合规")
		return nil
	}

	var account = Account{
		Number: number,
		Extra: extra,
		Password: password,
	}
	return &account
}


type NumHandle struct{

}

func (this *NumHandle) JudgeNum(num int) {
	if num % 2 == 0 {
		fmt.Println("是偶数")
	}else {
		fmt.Println("是奇数")
	}
}

func (this *NumHandle) Print(row int, col int, char string) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Print(char)
		}
		fmt.Println()
	}
}


type Calcuator struct{

}

func (this *Calcuator) Add(n1 float64, n2 float64) float64 {
	sum := n1 + n2
	fmt.Println("n1 + n2 =", sum)
	return sum
}

func (this *Calcuator) Sub(n1 float64, n2 float64) float64 {
	sub := n1 - n2
	fmt.Println("n1 - n2 =", sub)
	return sub
}

func (this *Calcuator) Mul(n1 float64, n2 float64) float64 {
	mul := n1 * n2
	fmt.Println("n1 * n2 =", mul)
	return mul
}

func (this *Calcuator) Div(n1 float64, n2 float64) float64 {
	div := n1 / n2
	fmt.Println("n1 ÷ n2 =", div)
	return div
}

func (this *Calcuator) Cal(n1 float64, n2 float64, op string) float64 {
	var res float64

	switch op{
	case "+":
		res = n1 + n2
		fmt.Println("n1 + n2 =", res)
	case "-":
		res = n1 - n2
		fmt.Println("n1 - n2 =", res)
	case "*":
		res = n1 * n2
		fmt.Println("n1 * n2 =", res)
	case "/":
		res = n1 / n2
		fmt.Println("n1 ÷ n2 =", res)
	default:
		fmt.Println("运算符有误")
	}
	return res
}



type MethodUtils struct{

}

//根据给定的数字打印乘法表
func (this *MethodUtils) MulTable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v x %v = %v\t", j, i, i * j)
		}
		fmt.Println()
	}
}

//转置一个3x3的二维数组
func (this *MethodUtils) TransArr(arr [3][3]int) {
	fmt.Println("原数组：")
	for _, v := range arr {
		for _, j := range v {
			fmt.Printf("%v\t", j)
		}
		fmt.Println()
	}
	fmt.Println("转置后数组：")
	for i := 0; i < 3; i++ {
		for _, v := range arr {
			fmt.Printf("%v\t", v[i])
		}
		fmt.Println()
	}

}
