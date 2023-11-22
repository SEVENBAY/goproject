package main
import (
	"fmt"
	"time"
)

func in() {
	for {
		fmt.Println("aaaaaaaaaaaaaaa")
		time.Sleep(time.Second)
	}
}

func out() {
	go in()
	
}

func main() {
	out()
	fmt.Println("出来了。。。")
	time.Sleep(time.Minute)
}

