package main
import (
	"fmt"
	"errors"
)

func test() {
	//错误处理机制
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("test()发生错误：", err)
		}
	}()

	var a int = 2
	var b int = 0
	c := a / b
	fmt.Println("c=", c)
}


func errorDef() {
	//错误处理
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("errorDef()发生错误:", err)
		}
	}()

	e := errors.New("这是一个自定义错误")
	panic(e)
}


func main()	{
	test()
	errorDef()
	fmt.Println("main()后边的代码")
}