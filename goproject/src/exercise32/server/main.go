package main
import (
	"fmt"
	"net"
	"encoding/binary"
	"bufio"
	"os"
)


func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:6564")
	if err != nil {
		fmt.Println("开启监听失败：", err)
		return
	}
	defer lis.Close()
	for {
		fmt.Println("开始等待客户端连接...")
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("等待客户端连接失败：", err)
			continue
		}
		fmt.Printf("与客户端<%v>建立连接...\n", conn.RemoteAddr())
		//发送欢迎语
		welcome := []byte("欢迎连接，我是小爱大脑")
		_, err = conn.Write(welcome)
		if err != nil {
			fmt.Println("发送欢迎信息失败")
			continue
		}
		for {
			//接收客户端发送数据的长度
			recv_dl := make([]byte, 4)
			_, err = conn.Read(recv_dl)
			if err != nil {
				fmt.Println("接收数据长度失败")
				continue
			}
			//接收数据
			data_len := binary.BigEndian.Uint32(recv_dl)
			recv_data := make([]byte, data_len)
			_, err = conn.Read(recv_data)
			if err != nil {
				fmt.Println("接收数据失败")
				continue
			}
			fmt.Printf("from %v：%v\n",conn.RemoteAddr(), string(recv_data))

			//向客户端发送数据
			fmt.Printf("to %v：", conn.RemoteAddr())
			reader := bufio.NewReader(os.Stdin)
			b_word, err := reader.ReadBytes('\n')
			if err != nil {
				fmt.Println("接收用户输入失败：", err)
				continue
			}
			//将长度转为byte
			send_dl := make([]byte, 4)
			binary.BigEndian.PutUint32(send_dl, uint32(len(b_word)))
			_, err = conn.Write(send_dl)
			if err != nil {
				fmt.Println("发送数据长度失败")
				continue
			}
			_, err = conn.Write(b_word)
			if err != nil {
				fmt.Println("发送数据失败")
				continue
			}
		}

	}

}