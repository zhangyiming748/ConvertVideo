package main

import (
	"github.com/zhangyiming748/ConvertVideo/constant"
	"github.com/zhangyiming748/ConvertVideo/conv"
	"github.com/zhangyiming748/ConvertVideo/util"
	"os"
)

func main() {
	if root := os.Getenv("root"); root == "" {
		constant.Root = "/data"
	} else {
		constant.Root = root
	}
	if to := os.Getenv("to"); to == "" {
		constant.To = "vp9"
	} else {
		constant.To = to
	}
	files := util.GetAllFiles(constant.Root)
	switch constant.To {
	case "vp9":
		for _, file := range files {
			conv.ProcessVideo2VP9(file)
		}
	default:
		os.Exit(0)
	}
}
