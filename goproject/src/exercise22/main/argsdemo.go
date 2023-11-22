package main
import (
	"fmt"
	"os"
	"flag"
)


type Config struct {
	host string
	port int
	username string
	password string
}


func getConfig() {
	var config Config
	flag.StringVar(&config.host, "h", "127.0.0.1", "主机地址")
	flag.IntVar(&config.port, "p", 3306, "主机端口")
	flag.StringVar(&config.username, "u", "", "用户名")
	flag.StringVar(&config.password, "pwd", "", "密码")
	flag.Parse()
	fmt.Printf("host=%v, port=%v, username=%v, password=%v\n", config.host, config.port, config.username, config.password)
}


func main() {
	for i, v := range os.Args {
		fmt.Printf("i=%v, v=%v\n", i, v)
	}
	getConfig()
}