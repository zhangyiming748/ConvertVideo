package main

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/constant"
	"github.com/zhangyiming748/ConvertVideo/conv"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/sql"
	"github.com/zhangyiming748/ConvertVideo/util"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if direction := os.Getenv("direction"); direction == "" {
		slog.Info("$direction为空,使用默认值", slog.String("$direction", constant.GetDirection()))
	} else {
		constant.SetDirection(direction)
		slog.Info("$direction不为空", slog.String("$direction", direction))
	}
	if root := os.Getenv("root"); root == "" {
		slog.Info("$root为空,使用默认值", slog.String("$root", constant.GetRoot()))
	} else {
		constant.SetRoot(root)
		slog.Info("$root不为空", slog.String("$root", root))
	}
	if to := os.Getenv("to"); to == "" {
		slog.Info("$to为空,使用默认值", slog.String("$to", constant.GetTo()))
	} else {
		constant.SetTo(to)
		slog.Info("$to不为空", slog.String("$to", to))
	}
	if level := os.Getenv("level"); level == "" {
		slog.Info("$level为空,使用默认值", slog.String("$level", constant.GetLevel()))
		setLog(constant.GetLevel())
	} else {
		constant.SetLevel(level)
		slog.Info("$level不为空", slog.String("$level", level))
		setLog(constant.GetLevel())
	}
	sql.SetEngine()

	err := filepath.Walk(constant.GetRoot(), func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			absPath, err := filepath.Abs(p)
			if err != nil {
				return err
			}
			fmt.Printf("准备处理的文件夹%v\n", info.Name())
			files := util.GetAllFiles(absPath)
			for _, file := range files {
				switch constant.To {
				case "vp9":
					conv.ProcessVideo2VP9(*mediainfo.GetBasicInfo(file))
				case "rotate":
					conv.RotateVideo(*mediainfo.GetBasicInfo(file), constant.GetDirection())
				case "merge":
					conv.MkvWithAss(*mediainfo.GetBasicInfo(file))
				default:
					os.Exit(0)
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	files := util.GetAllFiles(constant.Root)
	fmt.Printf("符合条件的文件:%v\n", files)

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
	fp := strings.Join([]string{constant.GetRoot(), "ConVideo.log"}, string(os.PathSeparator))
	fmt.Printf("数据库位置%v\n", fp)
	logf, err := os.OpenFile(fp, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}
