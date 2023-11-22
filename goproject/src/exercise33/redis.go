package main
import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)


func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("连接失败：", err)
		return
	}
	defer conn.Close()
	fmt.Println("连接成功！")
	//操作redis
	_, err = conn.Do("select", 2)
	if err != nil {
		fmt.Println("切换数据库失败：", err)
		return
	}
	_, err = conn.Do("set", "name", "jack")
	if err != nil {
		fmt.Println("设置值失败：", err)
		return
	}
	r, err := redis.ByteSlices(conn.Do("keys", "*"))
	if err != nil {
		fmt.Println("获取所有键失败：", err)
		return
	}
	fmt.Print("当前库中的键：")
	for _, v := range(r) {
		fmt.Print(string(v))
		fmt.Print("\t")
	}
	fmt.Println()
	value, err := redis.String(conn.Do("get", "name"))
	// value, err := conn.Do("get", "name")
	if err != nil {
		fmt.Println("设置值失败：", err)
		return
	}
	fmt.Printf("%v", value)

}