package main
import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"path/filepath"
	"strconv"
)


var sepLine string = "------------------------"


//展示某目录下文件列表
func listFile(dirname string) []string {
	//定义切片接收文件列表
	var file_list []string
	res, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println("读取目录失败，获取文件列表失败：", err)
		return file_list
	}
	fmt.Println(sepLine)
	for _, v := range res {
		if !v.IsDir() {
			file_list = append(file_list, v.Name())
			fmt.Println(v.Name())
		}
	}
	fmt.Println(sepLine)
	return file_list
}

//打印命令选项
func recvCommand() string {
	var command string
	fmt.Println("操作命令：")
	fmt.Print("\t1-查找\n\t2-替换\n\t3-增加编号\n\t4-增加前缀\n\t5-增加后缀\n\t0-退出\n\n")
	for {
		fmt.Print("请输入命令编号(0-5)：")
		fmt.Scanln(&command)
		switch command {
		case "1", "2", "3", "4", "5", "0":
			return command
		default:
			fmt.Println("输入的编号有误，请重新输入！")
			continue
		}
		// return command
	}
}

//查找功能
func findFiles(all_files []string, find_str string) []string {
	var files []string
	for _, v := range(all_files) {
		if strings.Contains(v, find_str) {
			files = append(files, v)
		}
	}
	return files
}

//替换功能
func replaceFiles(files []string, old_str string, new_str string, path string) {
	for _, f_name := range files {
		ext_n := filepath.Ext(f_name)
		name := f_name[:(len(f_name) - len(ext_n))]
		new_f_name := strings.Replace(name, old_str, new_str, -1)
		//拼接新文件名称
		new_f_name += ext_n
		//进行文件重命名
		err := os.Rename(filepath.Join(path, f_name), filepath.Join(path, new_f_name))
		if err != nil {
			fmt.Printf("文件%v重命名失败\n", f_name)
			fmt.Println(err)
		}
	}
}

//文件名称前增加编号
func addNum(files []string, path string) {
	for i, v := range files {
		new_name := strconv.Itoa(i + 1) + "_" + v
		//进行文件重命名
		err := os.Rename(filepath.Join(path, v), filepath.Join(path, new_name))
		if err != nil {
			fmt.Printf("文件%v重命名失败\n", v)
			fmt.Println(err)
		}
	}
}

//增加前缀
func addPrefix(files []string, prefix string, path string) {
	for _, v := range files {
		new_name := prefix + v
		//进行文件重命名
		err := os.Rename(filepath.Join(path, v), filepath.Join(path, new_name))
		if err != nil {
			fmt.Printf("文件%v重命名失败\n", v)
			fmt.Println(err)
		}
	}
}

//增加后缀
func addSuffix(files []string, suffix string, path string) {
	for _, f_name := range files {
		ext_n := filepath.Ext(f_name)
		name := f_name[:(len(f_name) - len(ext_n))]
		new_f_name := name + suffix
		//拼接新文件名称
		new_f_name += ext_n
		//进行文件重命名
		err := os.Rename(filepath.Join(path, f_name), filepath.Join(path, new_f_name))
		if err != nil {
			fmt.Printf("文件%v重命名失败\n", f_name)
			fmt.Println(err)
		}
	}
}


func main() {
	for {
		var path string

		fmt.Println()
		fmt.Print("请输入文件目录(目录名不能包含空格)(输入0退出)：")
		fmt.Scanln(&path)
		if path == "0" {
			fmt.Println("Bye bye...")
			return
		}
		//获取文件名称清单
		file_list := listFile(path)
		if len(file_list) == 0 {
			fmt.Println("该目录下未找到文件...")
			continue
		}
		//定义需要操作的文件集合
		op_files := file_list
		var end_loop bool = false
		for {
			//输出选项并接收用户输入
			command := recvCommand()
			switch command {
				case "1":
					for {
						var find_str string
						var do_op string
						fmt.Print("请输入要查找的内容(回车查询全部)：")
						fmt.Scanln(&find_str)
						if find_str != "" {
							op_tmp := op_files
							op_files = findFiles(op_files, find_str)
							if len(op_files) == 0 {
								fmt.Println("未查找到任何文件...")
								op_files = op_tmp
								continue
							}
						}
						// if find_str == "" {
						// 	op_files = file_list
						// } else {
						// 	op_files = findFiles(file_list, find_str)
						// 	if len(op_files) == 0 {
						// 		fmt.Println("未查找到任何文件...")
						// 		continue
						// 	}
						// }
						
						fmt.Println(sepLine)
						for _, v := range op_files {
							fmt.Println(v)
						}
						fmt.Println(sepLine)
						fmt.Print("是否针对查找到的文件进行操作？(y/n)：")
						fmt.Scanln(&do_op)
						if strings.ToLower(do_op) == "y" {
							break
						}
					}
				case "2":
					var old_s string
					var new_s string
					fmt.Print("请输入要被替换的内容：")
					fmt.Scanln(&old_s)
					if old_s == "" {
						fmt.Println("输入有误！")
						continue
					}
					fmt.Print("请输入要替换成的内容(回车替换成空)：")
					fmt.Scanln(&new_s)
					// if new_s == "" {
					// 	fmt.Println("输入有误！")
					// 	continue
					// }
					//对文件名称进行替换
					replaceFiles(op_files, old_s, new_s, path)
					listFile(path)
					end_loop = true
				case "3":
					addNum(op_files, path)
					listFile(path)
					end_loop = true
				case "4":
					var prefix string
					fmt.Print("请输入要增加的前缀：")
					fmt.Scanln(&prefix)
					if prefix == "" {
						fmt.Println("输入有误！")
						continue
					}
					addPrefix(op_files, prefix, path)
					listFile(path)
					end_loop = true
				case "5":
					var suffix string
					fmt.Print("请输入要增加的后缀：")
					fmt.Scanln(&suffix)
					if suffix == "" {
						fmt.Println("输入有误！")
						continue
					}
					addSuffix(op_files, suffix, path)
					listFile(path)
					end_loop = true
				case "0":
					fmt.Println("Bye bye...")
					return
			}
			if end_loop {
				break
			}
		}
	}
	
}