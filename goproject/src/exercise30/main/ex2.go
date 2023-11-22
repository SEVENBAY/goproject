package main
import (
	"fmt"
	"time"
)


func main() {
	var intChan = make(chan int, 1)
	go func () {
		for i := 0; i < 2; i++ {
			intChan <- i
			fmt.Printf("%d插入管道...\n", i)
			time.Sleep(time.Second * 1)
		}
		close(intChan)
	}()

	// for {
	// 	select {
	// 		case res := <-intChan:
	// 			fmt.Println("从管道获取数据：", res)
	// 		default:
	// 			continue
	// 			// fmt.Println("未获取到数据")
	// 	}
	// }
	for {
		res, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println(res, ok)
	}
	//只读或者只写管道
	var readOnlyChan <-chan int = make(chan int, 3)
	// r := <-readOnlyChan
	fmt.Println("readOnlyChan=", readOnlyChan)
	var writeOnlyChan chan<- int = make(chan int, 3)
	writeOnlyChan <- 10
	// r := <-writeOnlyChan
	fmt.Println("writeOnlyChan=", writeOnlyChan)
}