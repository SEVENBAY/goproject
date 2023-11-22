package main
import (
	"fmt"
	"time"
	"math/rand"
	"os"
	"strconv"
	"io/ioutil"
	"strings"

)

//向文件中写入1000个随机数据
func writeDataToFile(filepath string, end_chan *chan int) {
	//打开文件
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	defer file.Close()

	// rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		num := rand.Intn(10000)
		file.WriteString(strconv.Itoa(num) + "\n")
	}
	// time.Sleep(time.Second * 3)
	close(*end_chan)
}

func sort(src_file, dst_file string, end_chan *chan int, doneChan *chan int) {
	// for {
	// 	fmt.Println("开始等待文件...")
	// 	_, ok := <-*end_chan
	// 	if !ok {
	// 		fmt.Println("文件已准备好...")
	// 		break
	// 	}
	// }
	fmt.Println("开始等待文件...")
	<-*end_chan
	fmt.Println("文件已准备好，开始排序...")
	data_b, err := ioutil.ReadFile(src_file)
	
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	data_s := string(data_b)
	//去掉字符串前后空白内容
	data_s = strings.TrimSpace(data_s)
	data := strings.Split(data_s, "\n")
	//对data进行排序，冒泡排序
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data) - i - 1; j++ {
			left, _ := strconv.Atoi(data[j])
			right, _ := strconv.Atoi(data[j+1])
			if left > right {
				tmp := data[j]
				data[j] = data[j+1]
				data[j+1] = tmp
			}
		}
	}
	d_file, err := os.OpenFile(dst_file, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	defer d_file.Close()
	for _, v := range data {
		d_file.WriteString(v + "\n")
	}
	// time.Sleep(time.Second * 5)
	close(*doneChan)
}


func main() {
	start := time.Now().UnixNano()
	var eChan chan int = make(chan int, 1)
	var dChan = make(chan int, 1)
	go writeDataToFile("D:\\goproject\\static\\src_data.txt", &eChan)
	go sort("D:\\goproject\\static\\src_data.txt", "D:\\goproject\\static\\dst_data.txt", &eChan, &dChan)
	
	fmt.Println("等待全部完成...")
	<- dChan
	fmt.Println("全部执行完成...")
	fmt.Println("用时=", float64((time.Now().UnixNano() - start))/1000000000.0)
}