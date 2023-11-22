package main
import (
	"fmt"
	"scrnario_verfiy/src/theme"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"errors"
	"bufio"
	"golang.org/x/crypto/ssh"
	"strings"
	"os"
	"io/ioutil"
	"path"
)

var (
	win fyne.Window
	top_box *fyne.Container
	pwd_grid *fyne.Container
	key_grid *fyne.Container
	key_input_p *widget.Entry
	key_input_f *widget.Entry
	cmd_text *widget.Entry
	log_text *fyne.Container
	ip_addr *widget.Entry
	user_name *widget.Entry
	report_dir *widget.Entry
	connect_type *widget.RadioGroup
	check_system_type *widget.RadioGroup
	session *ssh.Session
	log_text_s *container.Scroll
)

//显示错误信息弹窗
func showErrorDialog(err error) {
	dialog.ShowError(err, win)
}

//连接方式选择回调函数
func connect_type_func(s string) {
	if s == "密码" {
		// key_input_f.SetText("")
		top_box.Remove(key_grid)
		top_box.Add(pwd_grid)
		top_box.Refresh()
	} else {
		// key_input_p.SetText("")
		top_box.Remove(pwd_grid)
		top_box.Add(key_grid)
		top_box.Refresh()
	}
}

//检查平台类型选择回调函数
func system_type_func(s string) {
	if s == "火天网境" {
		cmd_text.SetText("rm -rf /opt/fscr/fscr/verify_scenario_files&&docker exec -i fscr python manage.py verify_scenario -s ")
	} else {
		cmd_text.SetText("rm /opt/project/fsnr/verify_scenario.json&&docker exec -i fsnr python manage.py verify_scenario -s ")
	}
}

//ssh-key文件选择按钮回调函数
func ssh_file_select() {
	file_d := dialog.NewFileOpen(func(f fyne.URIReadCloser, err error) {
		if err != nil {
			showErrorDialog(err)
			return
		}
		if f != nil {
			key_input_f.SetText(f.URI().Path())
		}
	}, win)
	file_d.Show()
}

//entry校验回调函数
func NullValidate(s string) error {
	if s == "" {
		return Err("field")
	}
	return nil
}

func Err(msg string) error {
	message := fmt.Sprintf("%v 不能为空！", msg)
	return errors.New(message)
}

//命令执行
func cmdExec() {
	//判断当前是否存在执行命令的session
	if session != nil {
		dialog.ShowError(errors.New("当前正在执行命令，请勿重复执行"), win)
		return
	}
	c_type := connect_type.Selected
	//校验参数
	err := ip_addr.Validate()
	if err != nil {
		dialog.ShowError(Err("ip地址"), win)
		return
	}
	err = user_name.Validate()
	if err != nil {
		dialog.ShowError(Err("用户名"), win)
		return
	}	
	if c_type == "密码" {
		err = key_input_p.Validate()
		if err != nil {
			dialog.ShowError(Err("密码"), win)
			return
		}
	} else {
		err = key_input_f.Validate()
		if err != nil {
			dialog.ShowError(Err("ssh文件"), win)
			return
		}
	}
	
	err = cmd_text.Validate()
	if err != nil {
		dialog.ShowError(Err("命令"), win)
		return
	}
	err = report_dir.Validate()
	if err != nil {
		dialog.ShowError(Err("报告存放目录"), win)
		return
	}
	
	ip := ip_addr.Text+ ":22"
	username := user_name.Text
	key_p := key_input_p.Text
	key_f := key_input_f.Text
	cmd := cmd_text.Text
	report_path := report_dir.Text
	system_type := check_system_type.Selected
	//建立服务器连接
	var cli *Cli
	if c_type == "密码" {
		cli, err = NewCli(username, key_p, ip, key_f, true)
	} else {
		cli, err = NewCli(username, key_p, ip, key_f, false)
	}
	if err != nil {
		dialog.ShowError(err, win)
		return
	}
	//命令拼接参数
	args, err := readArgs()
	if err != nil {
		dialog.ShowError(err, win)
		return
	}
	if args == "" {
		dialog.ShowError(errors.New("参数文件<args.txt>为空！"), win)
		return
	}
	cmd = cmd + args
	if c_type == "密码" {
		//对命令进行加工
		cmd = fmt.Sprintf(`echo %v|sudo -S bash -c "%v"`, key_p, cmd)
	}
	//执行命令
	outReader, err := cli.RunCmdBuffer(cmd)
	if err != nil {
		dialog.ShowError(err, win)
		return
	}
	session = cli.session
	buf_reader := bufio.NewReader(outReader)
	go func() {
		log_text.RemoveAll()
		defer func() {
			if session != nil {
				session.Close()
				session = nil
			}
		}()
		for {
			content, err := buf_reader.ReadString('\n')
			if err != nil {
				break
			}
			content = strings.TrimSpace(content)
			log := canvas.NewText(content, color.White)
			log.TextSize = 10
			log_text.AddObject(log)
			log_text_s.ScrollToBottom()
		}
		//压缩并下载报告
		var srcPath, dest_file string
		if system_type == "火天网境" {
			compressCmd := "cd /opt/fscr/fscr&&tar zcf verify_scenario.tar.gz verify_scenario_files verify_scenario.json"
			if c_type == "密码" {
				compressCmd = fmt.Sprintf(`echo %v|sudo -S bash -c "%v"`, key_p, compressCmd)
			}
			_, err := cli.RunCmd(compressCmd)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			srcPath = "/opt/fscr/fscr/verify_scenario.tar.gz"
			dest_file = path.Join(report_path, "fscr_verify_scenario.tar.gz")
		} else {
			srcPath = "/opt/project/fsnr/verify_scenario.json"
			dest_file = path.Join(report_path, "fsnr_verify_scenario.json")
		}
		_, err = cli.ScpFileToLocal(srcPath, dest_file)
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		dialog.ShowInformation("报告下载完毕", "报告文件路径\n" + dest_file, win)
	}()
}

//停止执行命令
func cmdExecStop() {
	if session != nil {
		session.Close()
		session = nil
	}
}

//读取参数文件中的参数
func readArgs() (string, error) {
	file, err := os.Open("args.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()
	
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	//去除两端空格
	contentStr := strings.TrimSpace(string(content))
	contentArr := strings.Fields(contentStr)
	args := strings.Join(contentArr, ",")
	return args, nil
}

//报告存放目录选择回调函数
func reportDirSelect() {
	dir_d := dialog.NewFolderOpen(LocalDirFile, win)
	dir_d.Show()
}

func LocalDirFile(f fyne.ListableURI, err error) {
	if err != nil {
		dialog.ShowError(err, win)
		fmt.Println("目录选择失败，请重试！")
		return
	}
	if f != nil {
		contents, _ := f.List()
		report_dir.SetText(path.Dir(contents[0].Path()))
	}
}


func main() {
	m_app := app.New()
	//设置中文字体
	m_app.Settings().SetTheme(&theme.MyTheme{IsDark: true,})
	
	win = m_app.NewWindow("场景检查")
	win.Resize(fyne.NewSize(800, 500))
	win.SetFixedSize(true)
	
	//选择检查平台
	check_system := widget.NewLabel("检查平台")
	check_system_type = widget.NewRadioGroup([]string{"火天网境", "火天星云"}, system_type_func)
	check_system_type.Horizontal = true
	check_system_type.Required = true
	system_type_grid := container.NewAdaptiveGrid(3, container.NewHBox(layout.NewSpacer(), check_system), check_system_type)
	
	//服务器连接相关配置
	connect_way := widget.NewLabel("连接方式")
	connect_type = widget.NewRadioGroup([]string{"密码", "ssh文件"}, connect_type_func)
	connect_type.Horizontal = true
	connect_type.Required = true
	way_grid := container.NewAdaptiveGrid(3, container.NewHBox(layout.NewSpacer(), connect_way), connect_type)
	
	ip_label := widget.NewLabel("ip地址")
	ip_addr = widget.NewEntry()
	ip_addr.Validator = NullValidate
	ip_grid := container.NewAdaptiveGrid(3, container.NewHBox(layout.NewSpacer(), ip_label), ip_addr)
	
	user_label := widget.NewLabel("用户名")
	user_name = widget.NewEntry()
	user_name.Validator = NullValidate
	user_grid := container.NewAdaptiveGrid(3, container.NewHBox(layout.NewSpacer(), user_label), user_name)
	
	pwd_label := widget.NewLabel("密码")
	key_input_p = widget.NewEntry()
	key_input_p.SetPlaceHolder("密码")
	key_input_p.Validator = NullValidate
	pwd_grid = container.NewAdaptiveGrid(3, container.NewHBox(layout.NewSpacer(), pwd_label), key_input_p)
	
	ssh_label := widget.NewLabel("ssh文件")
	key_input_f = widget.NewEntry()
	key_input_f.SetPlaceHolder("ssh文件")
	key_input_f.Validator = NullValidate
	key_file_b := widget.NewButton("选择文件", ssh_file_select)
	key_grid = container.NewAdaptiveGrid(3, container.NewHBox(layout.NewSpacer(), ssh_label), key_input_f, container.NewAdaptiveGrid(3, key_file_b))
	
	top_box = container.NewVBox(system_type_grid, way_grid, ip_grid, user_grid)
	
	report_label := widget.NewLabel("报告存放目录")
	report_dir = widget.NewEntry()
	report_dir.Validator = NullValidate
	report_select := widget.NewButton("选择目录", reportDirSelect)
	report_dir_grid := container.NewAdaptiveGrid(3, container.NewHBox(layout.NewSpacer(), report_label), report_dir, container.NewAdaptiveGrid(3, report_select))
	
	//命令执行区域
	sep := widget.NewSeparator()
	cmd_label := canvas.NewText("执行命令", color.White)
	cmd_label.TextSize = 12
	
	cmd_text = widget.NewMultiLineEntry()
	cmd_text.Validator = NullValidate
	cmd_text.TextStyle.Bold = true
	cmd_text.SetMinRowsVisible(2)
	cmd_text.Disable()
	cmd_btn := widget.NewButton("执行", cmdExec)
	cmd_stop_btn := widget.NewButton("停止", cmdExecStop)
	btn_box := container.NewHBox(cmd_btn, cmd_stop_btn)
	m_border := container.NewBorder(cmd_label, nil, nil, btn_box, cmd_text)
	connect_type.SetSelected("ssh文件")
	check_system_type.SetSelected("火天网境")
	
	//执行日志显示区域
	log_label := canvas.NewText("执行日志", color.White)
	log_label.TextSize = 12
	log_text = container.NewVBox()
	log_text_s = container.NewScroll(log_text)
	log_text_con := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(800, 150)), log_text_s)
	
	
	middle_box := container.NewVBox(report_dir_grid, sep, m_border, log_label, log_text_con)
	
	my_border := container.NewBorder(top_box,nil, nil, nil, middle_box)
	
	win.SetContent(my_border)
	win.ShowAndRun()
}