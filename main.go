package main

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/constant"
	"github.com/zhangyiming748/ConvertVideo/conv"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/util"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

func init() {
	setLog()
}
func main() {
	go NumsOfGoroutine()
	go util.ExitAfterRun()
	t := new(util.ProcessDuration)
	t.SetStart(time.Now())
	defer func() {
		log.Printf("程序总用时:%v分\n", t.GetDuration().Minutes())
	}()

	if direction := os.Getenv("direction"); direction == "" {
		log.Printf("$direction为空,使用默认值%v\n", constant.GetDirection())
	} else {
		constant.SetDirection(direction)
		log.Printf("$direction不为空,修改为%v\n", constant.GetDirection())
	}
	if root := os.Getenv("root"); root == "" {
		log.Printf("$root为空,使用默认值%v\n", constant.GetRoot())
	} else {
		constant.SetRoot(root)
		log.Printf("$root不为空,修改为%v\n", constant.GetRoot())
	}
	if to := os.Getenv("to"); to == "" {
		log.Printf("$to为空,使用默认值%v\n", constant.GetTo())
	} else {
		constant.SetTo(to)
		log.Printf("$to不为空,修改为%v\n", constant.GetTo())
	}
	err := filepath.Walk(constant.GetRoot(), func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			absPath, err := filepath.Abs(p)
			if err != nil {
				return err
			}
			log.Printf("准备处理的文件夹%v\n", info.Name())
			files := util.GetAllFiles(absPath)
			cpus := constant.GetCpuNums()
			if cpus > constant.MaxCPU {
				cpus = constant.MaxCPU
			}
			var wg sync.WaitGroup
			ch := make(chan struct{}, cpus/4)
			log.Printf("CPU个数:%d\t协程缓冲区:%d\n", constant.GetCpuNums(), cpus/4)
			for _, file := range files {
				switch constant.To {
				case "vp9":
					ch <- struct{}{}
					go func() {
						wg.Add(1)
						conv.ProcessVideo2VP9(*mediainfo.GetBasicInfo(file))
						wg.Done()
						<-ch
					}()
				case "rotate":
					conv.RotateVideo(*mediainfo.GetBasicInfo(file), constant.GetDirection())
				case "merge":
					conv.MkvWithAss(*mediainfo.GetBasicInfo(file))
				case "clip":
					conv.ProcessVideo2clip(*mediainfo.GetBasicInfo(file))
				default:
					os.Exit(0)
				}
			}
			wg.Wait()
		}
		return nil
	})
	if err != nil {
		log.Println("Error:", err)
	}
	files := util.GetAllFiles(constant.Root)
	log.Printf("符合条件的文件:%v\n", files)
	t.SetEnd(time.Now())
}
func setLog() {
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   strings.Join([]string{constant.GetRoot(), "mylog.log"}, string(os.PathSeparator)),
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
	}

	// 创建一个用于输出到控制台的Logger实例
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)

	// 设置文件Logger
	//log.SetOutput(fileLogger)

	// 同时输出到文件和控制台
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 在这里开始记录日志

	// 记录更多日志...

	// 关闭日志文件
	//defer fileLogger.Close()
}
func NumsOfGoroutine() {
	for {
		fmt.Printf("当前程序运行时协程个数:%d\n", runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
}
