package main
import (
	"testing"
)


var mon = Monster{
	Name: "小明",
	Age: 23,
	Skill: "cb",
}
var file = `D:\goproject\static\monster_test.txt`


func TestStore(t *testing.T){
	mon.Store(file)
}

func TestReStore(t *testing.T) {
	m := mon.ReStore(file)
	_, ok := m.(Monster)
	if !ok {
		t.Fatal("ReStore测试不通过")
	}
}