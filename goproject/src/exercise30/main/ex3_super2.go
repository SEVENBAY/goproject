package main
import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"io/ioutil"
	"strings"
	"path/filepath"
	"time"

)

//向文件中写入1000个随机数据
func writeDataToFile(file_path string, fileNo int) {
	//打开文件
	file, err := os.OpenFile(file_path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	defer file.Close()

	for i := 0; i < 1000; i++ {
		num := rand.Intn(10000)
		file.WriteString(strconv.Itoa(num) + "\n")
	}
	noChan <- fileNo
}

func sort(src_filepath, dst_filepath string) {
	data_b, err := ioutil.ReadFile(src_filepath)
	
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
	d_file, err := os.OpenFile(dst_filepath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	defer d_file.Close()
	for _, v := range data {
		d_file.WriteString(v + "\n")
	}
	doneChan <- 1
}

var noChan = make(chan int, 10) //存放已经生成的文件编号
var doneChan = make(chan int, 10)

func main() {
	start := time.Now().UnixNano()
	src_file_name := "src_data.txt"
	dst_file_name := "dst_data.txt"

	//写10个文件
	for i := 0; i < 10; i++ {
		filePath := filepath.Join("D:\\goproject\\static", strconv.Itoa(i) + src_file_name)
		go writeDataToFile(filePath, i)
	}
	//随生成随排序
	fmt.Println("正在对文件执行排序操作...")
	for i := 0; i < 10; i++ {
		fileNo := <-noChan
		scrFile := filepath.Join("D:\\goproject\\static", strconv.Itoa(fileNo) + src_file_name)
		dstFile := filepath.Join("D:\\goproject\\static", strconv.Itoa(fileNo) + dst_file_name)
		go sort(scrFile, dstFile)
	}
	//等待排序完成
	for i := 0; i < 10; i++ {
		<-doneChan
	}
	fmt.Println("全部执行完成...")
	fmt.Println("用时=", float64((time.Now().UnixNano() - start))/1000000000.0)
}