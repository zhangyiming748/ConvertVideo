package constant

import (
	"runtime"
)

var (
	Root string = "/data" // 工作目录 如果为空  默认/data
	//To   string = "clip"                      // 转换到的编码 如果为空  默认vp9
	To string = "vp9" // 转换到的编码 如果为空  默认vp9
	//To        string = "merge" // 转换到的编码 如果为空  默认vp9
	Direction string = "ToRight"
	CpuNums   int    = runtime.NumCPU() // 核心数
)

const (
	MaxCPU = 12
)

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
