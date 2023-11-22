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
	if (this.tail + 1) % len(this.arr) == this.head {
		fmt.Println("队列已满")
		return -1
	}
	this.arr[this.tail] = data
	this.tail = (this.tail + 1) % len(this.arr)
	return 0
}


//从队列取数据
func(this *MyQue) Get() (bool, int) {
	if this.head == this.tail {
		// fmt.Println("队列为空")
		return false, -1
	}
	res := this.arr[this.head]
	this.head = (this.head + 1) % len(this.arr)
	// fmt.Println(res)
	return true, res
}


//展示队列中数据
func(this *MyQue) Show() {
	l_q := len(this.arr)
	if this.head == this.tail {
		fmt.Println("队列为空")
		return
	}
	//获取队列元素个数
	size := (this.tail + l_q - this.head) % l_q

	for i := this.head; i < this.head + size; i++ {
		fmt.Printf("%v ", this.arr[i % l_q])
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
	//启动消费协程
	go Consume(&my_q, 1)
	go Consume(&my_q, 2)
	go Consume(&my_q, 3)
	go Consume(&my_q, 4)
	go Consume(&my_q, 5)
	go Consume(&my_q, 6)

	no := 1
	for {
		lock.Lock()
		my_q.Put(no)
		lock.Unlock()
		no++
		time.Sleep(time.Second)
	}
}