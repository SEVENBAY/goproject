package main
import (
	"fmt"
)


type Animal struct{
	Sex string
	Age int
}


func (this *Animal) Eat() {
	fmt.Println("动物天生会吃饭:", this.Sex, this.Age)
}

func (this *Animal) Sleep() {
	fmt.Println("动物天生会睡觉:", this.Sex, this.Age)
}


type Cat struct{
	Name string
	Skill string
	Animal
}

func (this *Cat) Catch() {
	fmt.Printf("%v会抓老鼠\n", this.Name)
}

func (this *Cat) Sleep() {
	fmt.Printf("%v岁的%v要睡觉了\n", this.Age, this.Name)
}


func main() {
	var cat = Cat{
		Name: "咪咪",
		Skill: "抓老鼠",
		Animal: Animal{Sex: "雌性", Age: 3},
		
	}
	cat.Eat()
	cat.Sleep()
	cat.Animal.Sleep()
	cat.Catch()
}