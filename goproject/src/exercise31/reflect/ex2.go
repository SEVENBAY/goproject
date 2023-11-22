package main
import (
	"fmt"
	"reflect"
)


func main() {
	var v float64 = 1.2
	rt := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)
	fmt.Println("type=", rt)
	fmt.Println("kind=", rt.Kind())
	fmt.Println("value=", rv)
	//反射转为基本数据类型
	ri := rv.Interface()
	fmt.Printf("interface=%v, interface type=%T\n", ri, ri)
	var new float64 = ri.(float64)
	fmt.Printf("new=%v, new type=%T", new, new)
}