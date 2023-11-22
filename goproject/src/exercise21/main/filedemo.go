package main
import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func readFile() {
	file_path := "D:\\goproject\\static\\测试.txt"
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("文件打开错误：", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		data, err := reader.ReadString('\n')
		fmt.Print(string(data))
		if err == io.EOF {
			break
		}
	}
}


func writeFile() {
	file_path := "D:\\goproject\\static\\测试.txt"
	file, err := os.OpenFile(file_path, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败：", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("\n我是哈哈哈\n")
	writer.Flush()
}


func copyFile() {
	src_file_path := `D:\goproject\static\测试.txt`
	dst_file_path := `D:\goproject\static\测试1.txt`
	src, err := os.OpenFile(src_file_path, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开错误：", err)
		return
	}
	defer src.Close()
	src_reader := bufio.NewReader(src)

	dst, err := os.OpenFile(dst_file_path, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件错误：", err)
		return
	}
	defer dst.Close()
	dst_writer := bufio.NewWriter(dst)
	io.Copy(dst_writer, src_reader)
	fmt.Println("文件复制完毕！")
}

type fileCount struct {
	numCount int
	strCount int
	spaceCount int
	otherCount int
}

func fileCharCount() {
	var fileCount fileCount
	file_path := "D:\\goproject\\static\\测试.txt"
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("读取文件出错：", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		data, err := reader.ReadBytes('\n')
		for _, v := range data {
			switch {
			case v >= '0' && v <= '9':
				fileCount.numCount ++
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				fileCount.strCount ++
			case v == ' ':
				fileCount.spaceCount ++
			default:
				fileCount.otherCount ++
			}
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("数字个数=%v, 字符个数=%v, 空格个数=%v, 其他个数=%v\n", fileCount.numCount, fileCount.strCount, fileCount.spaceCount, fileCount.otherCount)

}


func main() {
	// readFile()
	// writeFile()
	// readFile()
	// copyFile()
	fileCharCount()
}