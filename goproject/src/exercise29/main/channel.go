package main
import (
	"fmt"
	"strconv"
)

/*
说明: 请完成如下案例
1) 创建一个 Person 结构体[Name, Age, Address]
2) 使用rand方法配合随机创建10个Person 实例，并放入到channel中
3) 遍历channel ，将各个Person实例的信息显示在终端...
*/

type Person struct{
	Name string
	Age int
	Address string
}



func main() {
	var channel = make(chan Person, 10)

	for i := 0; i < 10; i++ {
		var p = Person{
			Name: "小" + strconv.Itoa(i),
			Age: 10 + i,
			Address: "胡同" + strconv.Itoa(i),
		}
		channel <- p
	}
	close(channel)

	//遍历channel
	for v := range channel {
		fmt.Println(v)
	}
}