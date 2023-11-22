package main
import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxActive: 0,
		MaxIdle : 4,
		Dial: func()(redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}


func main() {
	conn := pool.Get()
	defer conn.Close()
	fmt.Println("获取连接成功！")
	//操作redis
	_, err := conn.Do("select", 2)
	if err != nil {
		fmt.Println("切换数据库失败：", err)
		return
	}
	_, err = conn.Do("set", "age", 30)
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
	value, err := redis.Int(conn.Do("get", "age"))
	// value, err := conn.Do("get", "name")
	if err != nil {
		fmt.Println("设置值失败：", err)
		return
	}
	fmt.Printf("%v", value)
	//关闭连接池
	pool.Close()
}