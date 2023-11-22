package main
import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//创建并初始化redis连接池
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxActive: 4,
		MaxIdle: 2,
		Dial: func()(redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func getKeys(conn redis.Conn, name string) [][]byte {
	res, err := redis.ByteSlices(conn.Do("keys", name))
	if err != nil {
		fmt.Println("获取key集合失败：", err)
		return [][]byte{}
	}
	fmt.Print("当前库中的key有：")
	for _, key := range res {
		fmt.Print(string(key))
		fmt.Print(" ")
	}
	fmt.Println()
	return res
}


func main() {
	//获取redis链接
	conn := pool.Get()
	defer conn.Close()
	//选择1库
	_, err := conn.Do("select", 1)
	if err != nil {
		fmt.Println("切换redis库失败：", err)
		return
	}
	
	for i := 0; i < 3; i++ {
		var name string
		var age int
		var skill string
		fmt.Printf("新建第%v个monster...\n", i + 1)
		fmt.Print("请输入name：")
		fmt.Scanln(&name)
		fmt.Print("请输入age：")
		fmt.Scanln(&age)
		fmt.Print("请输入skill：")
		fmt.Scanln(&skill)
		mon_key := fmt.Sprintf("monster_%v", i + 1)
		_, err = conn.Do("hset", mon_key, "name", name, "age", age, "skill", skill)
		if err != nil {
			fmt.Println("redis设置哈希失败：", err)
			return
		}
	}
	fmt.Println()
	all_key := getKeys(conn, "monster*")
	for _, k := range all_key {
		mon_key := string(k)
		name, err := redis.String(conn.Do("hget", mon_key, "name"))
		if err != nil {
			fmt.Printf("获取%v的name信息失败：%v\n", mon_key, err)
			continue
		}
		age, err := redis.Int(conn.Do("hget", mon_key, "age"))
		if err != nil {
			fmt.Printf("获取%v的age信息失败：%v\n", mon_key, err)
			continue
		}
		skill, err := redis.String(conn.Do("hget", mon_key, "skill"))
		if err != nil {
			fmt.Printf("获取%v的skill信息失败：%v\n", mon_key, err)
			continue
		}
		fmt.Printf("%v的信息如下：\n", string(k))
		fmt.Printf("name:%v\n", name)
		fmt.Printf("age:%v\n", age)
		fmt.Printf("skill:%v\n", skill)
	}
}