package main
import (
	"fmt"
	"reflect"
)


func main() {
	name := "Tom"
	t := reflect.TypeOf(name)
	v := reflect.ValueOf(name)
	fmt.Printf("type's type=%T, value's type=%T\n", t, v)
	fmt.Printf("type=%v, value=%v\n", t, v)
	fmt.Println("kind=", t.Kind())
	fmt.Println("name=", t.Name())
	fmt.Println("string=", t.String())
	// 进行值修改
	fmt.Println()
	var nickname = "Jerry"
	rv := reflect.ValueOf(&nickname)
	fmt.Println(rv.CanAddr())
	fmt.Println(rv.Elem().CanSet())
	rv.Elem().SetString("Jacky")
	fmt.Println("nickname=", nickname)
	//反射对象获取值
	fmt.Println()
	rv = reflect.ValueOf("TomandJerry")
	name = rv.Interface().(string)
	fmt.Printf("name=%v, type=%T\n", name, name)

}