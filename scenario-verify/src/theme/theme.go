package theme

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/theme"
    _"embed"
    "image/color"
)

//go:embed 1635950606516261.ttf
var MyFont []byte

var ResourceSourceHanSansTtf = &fyne.StaticResource{
	StaticName: "1635950606516261.ttf",
	StaticContent: MyFont,
}

type MyTheme struct{
	IsDark bool
}

var _ fyne.Theme = (*MyTheme)(nil)

// return bundled font resource
// ResourceSourceHanSansTtf 即是 bundle.go 文件中 var 的变量名
func (m MyTheme) Font(s fyne.TextStyle) fyne.Resource {
    return ResourceSourceHanSansTtf
}
func (m *MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	curTheme := theme.DefaultTheme()
	if !m.IsDark {
		curTheme = theme.LightTheme()
	}
    return curTheme.Color(n, v)
}

func (m *MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	curTheme := theme.DefaultTheme()
	if !m.IsDark {
		curTheme = theme.LightTheme()
	}
    return curTheme.Icon(n)
}

func (m *MyTheme) Size(n fyne.ThemeSizeName) float32 {
	curTheme := theme.DefaultTheme()
	if !m.IsDark {
		curTheme = theme.LightTheme()
	}
    return curTheme.Size(n)
}