package main
import (
	"fmt"
	"sync"
	"time"
)

//计算1-50的阶乘，将结果放入map中，最后遍历map

var (
	res = make(map[int]uint, 50)
	lock sync.Mutex
)

func unitCal(n int) {
	var total uint = 1
	for i := 1; i <= n; i++ {
		total *= uint(i)
	}
	lock.Lock()
	res[n] = total
	lock.Unlock()
}


func main() {
	for i := 1; i <= 50; i++ {
		go unitCal(i)
	}
	time.Sleep(time.Second * 5)
	//遍历map
	fmt.Println("map的长度=", len(res))
	lock.Lock()
	for i, v := range res{
		fmt.Printf("res[%v]=%v\n", i, v)
	}
	lock.Unlock()
}