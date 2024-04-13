package main

import (
	"github.com/zhangyiming748/ConvertVideo/constant"
	"github.com/zhangyiming748/ConvertVideo/conv"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/util"
	"io"
	"log/slog"
	"os"
)

func main() {
	if level := os.Getenv("level"); level == "" {
		constant.SetLevel("Debug")
		slog.Info("$level为空,使用默认值", slog.String("$root", constant.GetLevel()))
		setLog(constant.GetLevel())
	} else {
		constant.SetLevel(level)
		slog.Info("$level不为空", slog.String("$level", level))
		setLog(constant.GetLevel())
	}
	if root := os.Getenv("root"); root == "" {
		constant.SetRoot("/data")
		slog.Info("$root为空,使用默认值", slog.String("$root", constant.GetRoot()))

	} else {
		constant.SetRoot(root)
		slog.Info("$root不为空", slog.String("$root", root))
	}
	if to := os.Getenv("to"); to == "" {
		constant.SetTo("vp9")
		slog.Info("$to为空,使用默认值", slog.String("$to", constant.GetTo()))
	} else {
		constant.SetTo(to)
		slog.Info("$to不为空", slog.String("$to", to))
	}
	files := util.GetAllFiles(constant.Root)
	switch constant.To {
	case "vp9":
		for _, file := range files {
			conv.ProcessVideo2VP9(*mediainfo.GetBasicInfo(file))
		}
	default:
		os.Exit(0)
	}
}
func setLog(level string) {
	var opt slog.HandlerOptions

	switch level {
	case "Debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Debug("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	}
	file := "ConVideo.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0770)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}
