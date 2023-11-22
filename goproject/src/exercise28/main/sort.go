package main
import (
	"fmt"
	"sort"
)


type Hero struct{
	Name string
	Age int
	Hobby string
}

type HeroArr []Hero

func (this HeroArr) Len() int {
	return len(this)
}

func (this HeroArr) Less(i int, j int) bool {
	if this[i].Age > this[j].Age {
		return true
	} else {
		return false
	}
}

func (this HeroArr) Swap(i int, j int) {
	tmp := this[i]
	this[i] = this[j]
	this[j] = tmp
}


func main() {
	hero1 := Hero{
		Name: "孙悟空",
		Age: 10000,
		Hobby: "吃香蕉",
	}
	hero2 := Hero{
		Name: "猪八戒",
		Age: 5000,
		Hobby: "睡觉",
	}
	hero3 := Hero{
		Name: "沙悟净",
		Age: 2500,
		Hobby: "挑担子",
	}

	var data = HeroArr{hero2, hero1, hero3}
	fmt.Println(data)
	sort.Sort(data)
	fmt.Println(data)


}