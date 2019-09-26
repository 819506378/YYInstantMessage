package PrintScreen

import "fmt"

// 日志输出颜色
const (
	BLACK_COLOR       = iota // 黑色
	RED_COLOR                // 红色
	GREEN_COLOR              // 绿色
	YELLOW_COLOR             // 黄色
	BLUE_COLOR               // 蓝色
	PURPLE_COLOR             // 紫色
	ULTRAMARINE_COLOR        // 深蓝色
	WHITE_COLOR              // 白色
)

/* 背景色默认值 */
const SCREEN_BACKGROUND_COLOR = 40

/* 前景色默认值 */
const SCREEN_FOREGROUND_COLOR = 30

var (
	bPrintScreen    bool
	backGroundColor int = SCREEN_BACKGROUND_COLOR
	foreGroundColor int = SCREEN_FOREGROUND_COLOR
)

//SetLogPrintScreen 日志是否同时显示到终端
func SetLogPrintScreen(isPrintScreen bool) {
	bPrintScreen = isPrintScreen
}

//SetLogBackGroundColor 设置日志打印背景色
func SetLogBackGroundColor(color int) {
	if !bPrintScreen {
		return
	}

	backGroundColor = SCREEN_BACKGROUND_COLOR + color
}

//SetLogForeGroundColor 设置日志打印字体前景色
func SetLogForeGroundColor(color int) {
	if !bPrintScreen {
		return
	}

	foreGroundColor = SCREEN_FOREGROUND_COLOR + color
}

// 字体显示特效
const (
	Default   = iota // 默认终端显示
	HighLight        // 高亮显示
	UnderLine = 4    // 下划线显示
	Flicker          // 闪烁显示
	Inverse   = 7    // 反白显示
	Invisible        // 不可见
)

var specialEffect int

//SetLogFontSpecialEffect 设置字体打印特效
func SetLogFontSpecialEffect(effect int) {
	if !bPrintScreen {
		return
	}

	specialEffect = effect
}

func PrintScreen(log string) {
	fmt.Println(" %c[%d;%d;%dm%s%c[0m ", 0x1B, specialEffect, backGroundColor, foreGroundColor, log, 0x1B)
}
