package main
import (
	"testing"
	"fmt"
)


func TestArrSer(t *testing.T) {
	data := arrSer()
	if data == nil {
		t.Fatal("TestArrSer() 测试不通过")
	}
	fmt.Println("TestArrSer()测试通过")
}


func TestArrDer(t *testing.T) {
	data := arrDer([]byte{91, 49})
	if !data {
		t.Fatal("TestArrDer() 测试不通过")
	}
	fmt.Println("TestArrDer() 测试通过")
}