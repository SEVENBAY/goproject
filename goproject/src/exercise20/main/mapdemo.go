package main
import (
	"fmt"
)


func main() {
	// var m = make(map[int]string, 10)	
	m := map[int]string{1: "tom", 2: "jack"}
	m[3] = "marry"
	m[4] = "mike"
	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("key=%v value=%v\n", k, v)
	}
	//删除key
	delete(m, 1)
	fmt.Println(m)
	//map切片
	// var sm = make([]map[string]string, 5)
	
	// for i := 0; i < 3; i++ {
	// 	var name, gender, email string
	// 	var stu = make(map[string]string, 3)
	// 	fmt.Printf("请输入第%v个学生的信息\n", i + 1)
	// 	fmt.Print("请输入姓名：")
	// 	fmt.Scanln(&name)
	// 	fmt.Print("请输入性别：")
	// 	fmt.Scanln(&gender)
	// 	fmt.Print("请输入邮箱：")
	// 	fmt.Scanln(&email)
	// 	stu["name"] = name
	// 	stu["gender"] = gender
	// 	stu["email"] = email
	// 	sm[i] = stu
	// }
	// fmt.Println(sm)

	//定义通用类型的map
	var mmp = make(map[interface{}]interface{})
	mmp[1] = "tom"
	mmp["age"] = 10
	mmp["weight"] = 75.5
	mmp["male"] = true
	fmt.Println("新的mmp=", mmp)

}