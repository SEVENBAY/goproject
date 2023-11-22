package main
import (
	"fmt"
)



func main() {
	//十进制转其他进制
	var i int = 67
	fmt.Printf("十进制67对应的二进制是%b\n", i)  //1000011
	fmt.Printf("十进制67对应的八进制是%o\n", 0103)
	fmt.Printf("十进制67对应的十六进制是%x\n", 0x43)

	//其他进制转十进制
	var j int = 032
	var k int = 0x34
	fmt.Printf("二进制1001011转成十进制是%v\n", 75)
	fmt.Printf("八进制%v转成十进制是%v\n", j, 26)
	fmt.Printf("十六进制%v转成十进制是%v\n", k, 52)

	//逻辑与、或、异或，移位
	fmt.Printf("3&4=%v\n", (3 & 4))
	fmt.Printf("-3|4=%v\n", (-3 | 4))
	fmt.Printf("3^4=%v\n", (3 ^ 4))
	fmt.Printf("10>>4=%v\n", (10 >> 4))
	fmt.Printf("10<<4=%v\n", (10 << 4))
}