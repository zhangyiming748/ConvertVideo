package constant

import (
	"runtime"
	"strconv"
)

var (
	Root string = "/media/zen/swap/pikpak/wallpaper/telegram" // 工作目录 如果为空  默认/data
	To   string = "vp9"                                       // 转换到的编码 如果为空  默认vp9
	//To        string = "merge" // 转换到的编码 如果为空  默认vp9
	Direction      string = "ToRight"
	CpuNums        string
	MaxConcurrency int = 3
)

func GetMaxConcurrency() int {
	return MaxConcurrency
}
func GetCpuNums() string {
	return CpuNums
}
func SetCpuNums() {
	CpuNums = strconv.Itoa(runtime.NumCPU() / 2)
}
func GetDirection() string {
	return Direction
}
func SetDirection(s string) {
	Direction = s
}

func GetRoot() string {
	return Root
}
func SetRoot(s string) {
	Root = s
}

func GetTo() string {
	return To
}
func SetTo(s string) {
	To = s
}
