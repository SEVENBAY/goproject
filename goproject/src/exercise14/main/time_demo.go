package main
import (
	"fmt"
	"strconv"
	"time"
)


func test() {
	s := ""
	for i := 0; i < 100000; i++ {
		s += "hello" + strconv.Itoa(i)
	}
}


func main() {
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("(go)耗时=%vs\n", end - start)
}