package main
import (
	"fmt"
)


func main() {
	var i float32 = 100
	var j float32 = 23.432
	fmt.Println("i+j=", i + j)
	fmt.Println("i*j=", i * j)
	fmt.Println("i/j=", i / j)
	i++
	fmt.Println("i=", i)
}