package main
import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
	"errors"
)

type Cli struct {
	username string
	password string
	host string
	key_path string
	client *ssh.Client
	session *ssh.Session
}

//建立连接
func NewCli(username, password, host, key_path string, by_pwd bool) (*Cli, error) {
	cli := &Cli{
		username: username,
		password: password,
		host: host,
		key_path: key_path,
	}
	//区分不同连接方式
	var auth_method ssh.AuthMethod
	if by_pwd {
		auth_method = ssh.Password(cli.password)
	} else {
		auth_method = func() ssh.AuthMethod {
			key, err := ioutil.ReadFile(key_path)
			if err != nil {
				return nil
			}
			signer, err := ssh.ParsePrivateKey(key)
			if err != nil {
				return nil
			}
			return ssh.PublicKeys(signer)
		}()
	}
	if auth_method == nil {
		return nil, errors.New("login failed")
	}

	//连接配置初始化
	con_config := &ssh.ClientConfig{
		User: cli.username,
		Auth: []ssh.AuthMethod{
			auth_method,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", cli.host, con_config)
	if err != nil {
		return nil, err
	}
	cli.client = client
	return cli, nil
}

//执行命令
func (cli *Cli) RunCmdBuffer(cmd string) (io.Reader, error) {
	session, err := cli.client.NewSession()
	if err != nil {
		return nil, err
	}
	cli.session = session
	
	outReader, err := session.StdoutPipe()
	if err != nil {
		return nil, err
	}
	go func() {
		err = session.Run(cmd)
		if err != nil {
			return
		}
	}()
	return outReader, nil
}

//执行命令
func (cli *Cli) RunCmd(cmd string) (string, error) {
	session, err := cli.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	out, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

//拷贝文件：从远程拷贝到本地
func (cli *Cli) ScpFileToLocal(src, dest string) (int64, error) {
	sftpClient, err := sftp.NewClient(cli.client)
	if err != nil {
		return 0, err
	}
	defer sftpClient.Close()

	src_file, err := sftpClient.Open(src)
	if err != nil {
		return 0, err
	}
	defer src_file.Close()
	dest_file, err := os.OpenFile(dest, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0666)
	if err != nil {
		return 0, err
	}
	defer dest_file.Close()

	n, err := io.Copy(dest_file, src_file)
	if err != nil {
		return 0, err
	}
	return n, nil
}

//拷贝文件：从本地拷贝到远程
func (cli *Cli) ScpFileToRemote(src, dest string) (int64, error) {
	sftpClient, err := sftp.NewClient(cli.client)
	if err != nil {
		return 0, err
	}
	defer sftpClient.Close()

	dest_file, err := sftpClient.Create(dest)
	if err != nil {
		return 0, err
	}
	defer dest_file.Close()
	src_file, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer src_file.Close()

	n, err := io.Copy(dest_file, src_file)
	if err != nil {
		return 0, err
	}
	return n, nil
}

//拷贝进度：从本地到远程
func (cli *Cli) ProcessToRemote(src, dest string) (float64, error) {
	//获取本地文件的总大小
	fileinfo, err := os.Stat(src)
	if err != nil {
		return 0.0, err
	}
	totalSize := fileinfo.Size()

	//获取远程文件大小
	cmd := fmt.Sprintf("du -b %v", dest)
	out, err := cli.RunCmd(cmd)
	if err != nil {
		return 0.0, err
	}
	size := handleSizeStr(out)
	return size / float64(totalSize), nil
}

//拷贝进度：从远程到本地
func (cli *Cli) ProcessToLocal(src, dest string) (float64, error) {
	//获取本地文件的总大小
	fileinfo, err := os.Stat(dest)
	if err != nil {
		return 0.0, err
	}
	curSize := fileinfo.Size()

	//获取远程文件总大小
	cmd := fmt.Sprintf("du -b %v", src)
	out, err := cli.RunCmd(cmd)
	if err != nil {
		return 0.0, err
	}
	totalSize := handleSizeStr(out)
	return float64(curSize) / totalSize, nil
}

//处理远程返回的size字符串
func handleSizeStr(s string) float64 {
	res := strings.Fields(s)
	size, _ := strconv.Atoi(res[0])
	return float64(size)
}
