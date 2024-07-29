package constant

import (
	"runtime"
)

var (
	Root       string = "/data" // 工作目录 如果为空  默认/data
	To         string = "h265"  // 转换到的编码 如果为空  默认vp9
	Direction  string = "ToRight"
	CpuNums    int    = runtime.NumCPU() // 核心数
	TransTitle bool
)

const (
	MaxCPU = 12
)

func SetTransTitle(s string) {
	if s == "0" {
		TransTitle = false
	} else {
		TransTitle = true
	}
}
func GetTransTitle() bool {
	return TransTitle
}
func GetCpuNums() int {
	return CpuNums
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
