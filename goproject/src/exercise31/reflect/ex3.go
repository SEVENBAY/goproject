package main
import (
	"fmt"
	"reflect"
)


type Cal struct{
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

func (this *Cal) GetSub(name string) int {
	sub := this.Num1 - this.Num2
	fmt.Printf("%v完成了减法运算，%v-%v=%v\n", name, this.Num1, this.Num2, sub)
	return sub
}


func main() {
	cal := Cal{
		Num1: 8,
		Num2: 3,
	}

	rt := reflect.TypeOf(cal)
	rv := reflect.ValueOf(cal)
	fmt.Println("rt=", rt)
	fmt.Println("rv=", rv)
	fieldNum := rt.NumField()
	fmt.Println("fieldNum=", fieldNum)
	//遍历字段
	for i := 0; i < fieldNum; i++ {
		fmt.Println("==================")
		field := rt.Field(i)
		fmt.Println("字段Name=", field.Name)
		fmt.Println("字段Type=", field.Type)
		fmt.Println("字段Tag=", field.Tag)
		field_v := rv.FieldByName(field.Name).Int()
		fmt.Println("字段Value=", field_v)
	}

	//调用函数
	fmt.Println()
	rt = reflect.TypeOf(&cal)
	rv = reflect.ValueOf(&cal)
	methodNum1 := rt.NumMethod()
	methodNum2 := rv.NumMethod()
	fmt.Printf("methodNum1=%d, methodNum2=%d\n", methodNum1, methodNum2)
	for i := 0; i < methodNum1; i++ {
		method := rt.Method(i)
		fmt.Println("方法Name=", method.Name)
		fmt.Println("方法Type=", method.Type)
		fmt.Println("方法Func=", method.Func)
		fmt.Println("方法Numin=", method.Type.NumIn())
		fmt.Println("方法Numin0=", method.Type.In(0))
		fmt.Println("方法Numin1=", method.Type.In(1))
		fmt.Println("方法Numout=", method.Type.NumOut())
		fmt.Println("方法Numout0=", method.Type.Out(0))

	}
	fmt.Println()
	method := rv.MethodByName("GetSub")
	res := method.Call([]reflect.Value{reflect.ValueOf("小红红")})
	fmt.Println("res=", res[0].Int())
}