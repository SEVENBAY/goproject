package main

import (
	"timing_remind/src/theme"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/dialog"
	"time"
	"image/color"
	"strconv"
	"os/exec"
	"strings"
)

var (
	nextTime *canvas.Text
	timeSep *widget.Entry
	msgText *widget.Entry
	myWin fyne.Window
	myApp fyne.App
	remindTime string
	remindSepTime int
	remindMsg string
	beginLoop bool
	remindType string
)

func addTrayIcon(a fyne.App, w fyne.Window) {
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	
	if desk, ok := a.(desktop.App); ok {
		me := fyne.NewMenu("menu", fyne.NewMenuItem("Open", func() {
			w.Show()
		}))
		desk.SetSystemTrayMenu(me)
	}
}

//设置按钮回调函数
func setBtn() {
	time_sep := timeSep.Text
	msg := msgText.Text
	
	if time_sep == "" {
		dialog.ShowInformation("提示","请填写间隔时间！", myWin)
		return
	}
	if msg == "" {
		dialog.ShowInformation("提示", "请填写提示信息！", myWin)
		return
	}
	
	//将时间间隔进行转换
	time_sep_i, err := strconv.Atoi(time_sep)
	if err != nil {
		dialog.ShowInformation("提示", "时间间隔只能是整数！", myWin)
		return
	}
	if time_sep_i <= 0 {
		dialog.ShowInformation("提示", "请填写大于0的数字！", myWin)
		return
	}
	remindMsg = msg
	remindSepTime = time_sep_i
	remindTime = time.Now().Add(time.Duration(remindSepTime) * time.Minute).Format("2006-01-02 15:04:05")
	nextTime.Text = remindTime
	nextTime.Refresh()
	beginLoop = true
}

//提醒定时器
func remindTimer() {
	for {
		if beginLoop {
			nowT := time.Now().Format("2006-01-02 15:04:05")
			comp := strings.Compare(remindTime, nowT)
			//当提醒时间比当前时间小时，则重置提醒时间
			if comp == -1 {
				remindTime = time.Now().Add(time.Duration(remindSepTime) * time.Minute).Format("2006-01-02 15:04:05")
				nextTime.Text = remindTime
				nextTime.Refresh()
			}
			if comp == 0 {
				if remindType == "锁屏" {
					lockCmd := exec.Command("rundll32.exe", "user32.dll", "LockWorkStation")
					lockCmd.Start()
				} else {
					noti := fyne.NewNotification("提醒消息", remindMsg)
					myApp.SendNotification(noti)
				}
				remindTime = time.Now().Add(time.Duration(remindSepTime) * time.Minute).Format("2006-01-02 15:04:05")
				nextTime.Text = remindTime
				nextTime.Refresh()
				go remindTimer()
				break
			}
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func init() {
	//启动定时器
	go remindTimer()
}

func main() {
	myApp = app.NewWithID("定时提醒")
	//添加系统图标
	icon, _ := fyne.LoadResourceFromPath("clock.jpeg")
	myApp.SetIcon(icon)
	//设置主题
	myApp.Settings().SetTheme(&theme.MyTheme{IsDark:true,})
	myWin = myApp.NewWindow("定时提醒")
	//设置窗口大小
	myWin.Resize(fyne.NewSize(400, 300))
	//设置系统托盘图标
	addTrayIcon(myApp, myWin)
	
	//显示当前实时时间
	curTime := canvas.NewText("", color.RGBA{255, 0, 0, 255})
	curTime.TextSize = 25
	go func() {
		for {
			ti := time.Now().Format("2006-01-02 15:04:05")
			curTime.Text = ti
			curTime.Refresh()
			time.Sleep(time.Second)
		}
	}()
	timeBorder := container.NewBorder(nil, widget.NewLabel(""), nil, nil, container.NewCenter(curTime))
	
	//提醒方式
	typeLabel := widget.NewLabel("提醒方式")
	typeRadio := widget.NewRadioGroup([]string{"消息", "锁屏"}, func(s string) {
		remindType = s
	})
	typeRadio.Required = true
	typeRadio.Horizontal = true
	typeRadio.SetSelected("消息")
	typeGrid := container.NewAdaptiveGrid(3, container.NewHBox(layout.NewSpacer(), typeLabel), typeRadio)
	
	//时间间隔设置
	timeSepLabel := widget.NewLabel("时间间隔")
	timeSep = widget.NewEntry()
	timeSep.SetText("30")
	timeSepSuf := widget.NewLabel("分钟")
	timeSepGrid := container.NewAdaptiveGrid(3, container.NewHBox(layout.NewSpacer(), timeSepLabel), timeSep, timeSepSuf)
	//提醒消息
	msgLabel := widget.NewLabel("提醒信息")
	msgText = widget.NewEntry()
	msgText.SetText("起来活动活动老腰吧亲！眼也受不了了！")
	msgGrid := container.NewAdaptiveGrid(3, container.NewHBox(layout.NewSpacer(), msgLabel), msgText)
	//确认按钮
	btn := widget.NewButton("开启提醒", setBtn)
	btnGrid := container.NewAdaptiveGrid(3, layout.NewSpacer(), container.NewAdaptiveGrid(3, layout.NewSpacer(), btn))
	
	//设置结果显示
	resLabel := widget.NewLabel("下一次提醒时间：")
	nextTime = canvas.NewText("", color.RGBA{255, 0, 0, 255})
	resBox := container.NewHBox(resLabel, nextTime)
	
	vBox := container.NewVBox(timeBorder, typeGrid, timeSepGrid, msgGrid, btnGrid, resBox)
	myWin.SetContent(vBox)
	myWin.ShowAndRun()
}