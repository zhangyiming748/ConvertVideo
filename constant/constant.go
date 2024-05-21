package constant

import (
	"runtime"
	"strconv"
)

var (
	Root      string = "/mnt/e/pikpak/hidden" // 工作目录 如果为空  默认/data
	To        string = "vp9"                  // 转换到的编码 如果为空  默认vp9
	Level     string = "Debug"                //日志的输出等级
	Direction string = "ToRight"
	CpuNums   string
)

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
func GetLevel() string {
	return Level
}
func SetLevel(s string) {
	Level = s
}
