package main
import (
	"fmt"
)

func genNum(nc chan int) {
	for i := 1; i <= 200000; i++ {
		nc<- i
	}
	close(nc)
}

func calNum(nc chan int, rc chan map[int]int, ec chan int) {
	for v := range nc {
		sum := 0
		for i := 1; i <= v; i++ {
			sum += i
		}
		var res = make(map[int]int, 1)
		res[v] = sum
		rc <- res
	}
	ec <- 1
}

func main() {
	var numChan = make(chan int, 2000000)
	var resChan = make(chan map[int]int, 2000000)
	var endChan = make(chan int, 16)

	//生成数据
	go genNum(numChan)
	//启动8个协程计算
	for i := 0; i < 16; i++ {
		go calNum(numChan, resChan, endChan)
	}

	//等待8个协程完毕
	for i := 0; i < 16; i++ {
		<-endChan
	}
	close(resChan)

	resChanLen := len(resChan)
	//遍历结果
	for v := range resChan {
		for i, j := range v {
			fmt.Printf("res[%d]=%d\n", i, j)
		}
	}
	fmt.Println("resChan的长度=", resChanLen)

}