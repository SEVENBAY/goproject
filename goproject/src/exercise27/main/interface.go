package main
import (
	"fmt"
)

type usb interface{
	load() //加载usb
	unload() //卸载usb
}

type computer struct{
	name string
	price float64
}

func (this *computer) load() {
	fmt.Printf("<%v>加载usb...\n", this.name)
}

func (this *computer) unload() {
	fmt.Printf("<%v>卸载usb...\n", this.name)
}

type phone struct{
	name string
}

func (this *phone) load() {
	fmt.Printf("<%v>加载usb...\n", this.name)
}

func (this *phone) unload() {
	fmt.Printf("<%v>卸载usb...\n", this.name)
}


func main() {
	var u usb
	u = &computer{
		name: "联想",
		price: 3005,
	}
	u.load()
	u.unload()

	u = &phone{
		name: "华为手机",
	}
	u.load()
	u.unload()

	new, ok := u.(*phone)
	if ok {
		fmt.Printf("%v", new)
	} else {
		fmt.Println("u不是phone")
	}
}

