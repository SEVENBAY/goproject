package main
import (
	"fmt"
	"strings"
)


func makeSuffix(suf string) func(string) string {
	checkSuffix := func(name string) string {
		if !strings.HasSuffix(name, suf) {
			name += suf
		}
		return name
	}
	return checkSuffix
}


func main() {
	f := makeSuffix(".jpg")
	res := f("winter")
	res1 := f("summar.jpg")
	fmt.Println("res=", res)
	fmt.Println("res1=", res1)
}