package main
import (
	"fmt"
	"net"
	"encoding/binary"
	"bufio"
	"os"
)


func main() {
	//连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:6564")
	if err != nil {
		fmt.Println("连接服务端失败")
		return
	}
	defer conn.Close()
	fmt.Println("连接到服务端：", conn.RemoteAddr())
	//接收欢迎信息
	welcome := make([]byte, 100)
	_, err = conn.Read(welcome)
	if err != nil {
		fmt.Println("欢迎信息接收失败")
		return
	}
	fmt.Printf("from %v：%v\n", conn.RemoteAddr(), string(welcome))

	for {
		fmt.Printf("to %v：", conn.RemoteAddr())
		reader := bufio.NewReader(os.Stdin)
		b_word, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("接收用户输入失败：", err)
			return
		}
		//向服务端发送数据长度
		data_len := make([]byte, 4)
		binary.BigEndian.PutUint32(data_len, uint32(len(b_word)))
		_, err = conn.Write(data_len)
		if err != nil {
			fmt.Println("向服务端发送数据长度失败：", err)
			return
		}
		_, err = conn.Write(b_word)
		if err != nil {
			fmt.Println("向服务端发送数据失败：", err)
			return
		}
		//接收服务器发送数据的长度
		recv_dl := make([]byte, 4)
		_, err = conn.Read(recv_dl)
		if err != nil {
			fmt.Println("接收服务端数据长度失败：", err)
			return
		}
		recv_data_len := binary.BigEndian.Uint32(recv_dl)
		recv_data := make([]byte, recv_data_len)
		_, err = conn. Read(recv_data)
		if err != nil {
			fmt.Println("接收服务端数据失败：", err)
			return
		}
		fmt.Printf("from %v：%v\n", conn.RemoteAddr(), string(recv_data))

	}
}