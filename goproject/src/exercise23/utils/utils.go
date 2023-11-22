package main
import (
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
)

type Monster struct{
	Name string
	Age int
	Skill string
}


func (this *Monster) Store(filePath string) {
	//进行序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("序列化失败：", err)
		return
	}

	//写入文件
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}
	defer file.Close()
	file.Write(data)
}

func (this *Monster) ReStore(filePath string) interface{} {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return nil
	}
	var m Monster
	err = json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println("发序列化失败：", err)
		return nil
	}
	return m
}

func main() {
	var monster = Monster{
		Name: "小红",
		Age: 22,
		Skill: "fs",
	}
	file := `D:\goproject\static\monster.txt`
	monster.Store(file)
	monster.ReStore(file)
}