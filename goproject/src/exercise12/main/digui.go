package main
import (
	"fmt"
)


/*
有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个!以后每天猴子都吃其中的一半，然后再多吃一个。
当到第十天时，想再吃时 (还没吃)，发现只有1个桃子了。问题:最初共多少个桃子?
*/
func eatP(days int) uint {
	if days == 1 {
		return 1
	}
	days--
	return (eatP(days) + 1) * 2
}




func main() {
	sum := eatP(10)
	fmt.Println("最初的桃子个数为", sum)
}