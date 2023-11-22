package main
import (
	"fmt"
	"goproject/src/exercise26/model"
)


func main() {
	account := model.NewAccount("120110", 119.9, "123432")
	if account == nil {
		return
	}
	fmt.Printf("account=%p, account的类型=%T\n", account, account)
	account.ShowDetail()
	//重新设置值
	account.SetExtra(21)
	account.ShowDetail()

	fmt.Println("===================")
	num_handle := model.NumHandle{}
	num_handle.JudgeNum(4)
	num_handle.JudgeNum(5)
	num_handle.Print(10, 50, "&")

	fmt.Println("===================")
	cal := model.Calcuator{}
	cal.Add(2.2, 3.3)
	cal.Cal(2.2, 3.3, "+")
	cal.Cal(2.2, 3.3, "*")
	cal.Cal(2.2, 3.3, "/")
	cal.Cal(2.2, 3.3, "&")

	fmt.Println("===================")
	uti := model.MethodUtils{}
	uti.MulTable(7)
	var arr = [3][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}, [3]int{7, 8, 9}}
	uti.TransArr(arr)
}