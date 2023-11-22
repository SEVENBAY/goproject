package main
import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var lock = &sync.Mutex{}

//数组模拟队列
type MyQue struct{
	arr [10]int
	tail int //数据入队列指针
	head int //数据出队列指针
}

//向队列中放数据
func(this *MyQue) Put(data int) int {
	//如果尾部指针达到数组长度，证明队列已满
	if this.tail == len(this.arr) {
		fmt.Println("队列已满")
		return -1
	}
	this.arr[this.tail] = data
	this.tail++
	return 0
}


//从队列取数据
func(this *MyQue) Get() (bool, int) {
	if this.head == this.tail {
		// fmt.Println("队列为空")
		return false, -1
	}
	res := this.arr[this.head]
	this.head++
	// fmt.Println(res)
	return true, res
}


//展示队列中数据
func(this *MyQue) Show() {
	if this.head == this.tail {
		fmt.Println("队列为空")
		return
	}
	for i := this.head; i < this.tail; i++ {
		fmt.Printf("%v ", this.arr[i])
	}
	fmt.Println()
}

//消费者
func Consume(que *MyQue, no int) {
	for {
		lock.Lock()
		status, data := que.Get()
		lock.Unlock()
		if status {
			fmt.Printf("%v号协程服务-->%v号客户\n", no, data)
		}
		time.Sleep(time.Second * time.Duration(rand.Intn(3)))
		// time.Sleep(time.Second)
	}
}


func main() {
	var my_q MyQue
	//启动消费者
	go Consume(&my_q, 1)
	go Consume(&my_q, 2)
	go Consume(&my_q, 3)

	//向队列放数据
	for i := 0; i < 10; i++ {
		lock.Lock()
		my_q.Put(i+1)
		lock.Unlock()
		time.Sleep(time.Second * time.Duration(rand.Intn(3)))
		// time.Sleep(time.Second)
	}
	//等待队列全部消费完
	for {
		time.Sleep(time.Second)
		if my_q.tail == 10 && my_q.head == 10 {
			break
		}
	}

}