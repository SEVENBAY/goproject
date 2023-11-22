package main
import (
	"fmt"
	"encoding/json"
)


//数组序列化
func arrSer() []byte {
	var myArr = [5]int{1, 2, 3, 4, 5}
	data, err := json.Marshal(myArr)
	if err != nil {
		fmt.Println("myArr序列化失败：", err)
		return nil
	}
	fmt.Printf("myArr序列化=%v\n", string(data))
	return data
}

//数组反序列化
func arrDer(data []byte) bool {
	var arr [5]int
	err := json.Unmarshal(data, &arr)
	if err != nil {
		fmt.Println("arr反序列化失败：", err)
		return false
	}
	fmt.Println("arr=", arr)
	return true
}


//map序列化
func mapSer() []byte {
	var myMap = map[string]interface{}{
		"name": "tom",
		"age": 20,
		"male": true,
	}
	data, err := json.Marshal(myMap)
	if err != nil {
		fmt.Println("myMap序列化失败：", err)
		return nil
	}
	fmt.Println("myMap序列化=", string(data))
	return data
}

//map反序列化
func mapDer(data []byte) {
	var mmap map[string]interface{}
	err := json.Unmarshal(data, &mmap)
	if err != nil {
		fmt.Println("map反序列化失败：", err)
		return
	}
	fmt.Println("mmap=", mmap)
}



//结构体序列化
func structSer() []byte {
	type Stru struct {
		Name string `json:"name"`
		Age int `json:"age"`
		Skill string `json:"skill"`
	}

	myStru := Stru{"jack", 22, "铁砂掌"}
	data, err := json.Marshal(myStru)
	if err != nil {
		fmt.Println("myStru序列化失败：", err)
		return nil
	}
	fmt.Println("myStru序列化=", string(data))
	return data
}

//结构体反序列化
func structDer(data []byte) {
	type s struct {
		Name string
		Age int
		Skill string
	}

	var ms s
	err := json.Unmarshal(data, &ms)
	if err != nil {
		fmt.Println("struct反序列化失败：", err)
		return
	}
	fmt.Println("ms=", ms)
}


func main() {
	a := arrSer()
	arrDer(a)
	m := mapSer()
	mapDer(m)
	s := structSer()
	structDer(s)
	
}