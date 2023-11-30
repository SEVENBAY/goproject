package main
import (
	"fmt"
	_"strings"
	_"bufio"
	"os"
	"io/ioutil"
)


func main() {
	s := "E:\\archer.wiki"
	res, err := ioutil.ReadDir(s)
	if err != nil {
		fmt.Println("读取目录失败，获取文件列表失败：", err)
		return
	}
	for _, v := range res {
		if !v.IsDir() {
			fmt.Println(v.Name())
		}
	}
	r1 := "E:\\archer.wiki\\aaa\\aa.txt"
	r2 := "E:\\archer.wiki\\aaa\\lc_aa.txt"
	os.Rename(r2, r1)
}