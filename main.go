package main

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/constant"
	"github.com/zhangyiming748/ConvertVideo/conv"
	mylog "github.com/zhangyiming748/ConvertVideo/log"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/util"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
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
	mylog.SetLog()
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
			for _, file := range files {
				switch constant.To {
				case "vp9":
					conv.ProcessVideo2VP9(*mediainfo.GetBasicInfo(file))
				case "h265":
					conv.ProcessVideo2H265(*mediainfo.GetBasicInfo(file))
				case "rotate":
					conv.RotateVideo(*mediainfo.GetBasicInfo(file), constant.GetDirection())
				case "merge":
					conv.MkvWithAss(*mediainfo.GetBasicInfo(file))
				case "clip":
					conv.ProcessVideo2clip(*mediainfo.GetBasicInfo(file))
				default:
					log.Fatalf("$to=%v参数错误\n", constant.GetTo())
				}
			}
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

func NumsOfGoroutine() {
	for {
		fmt.Printf("\r当前程序运行时协程个数:%d\n", runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
}
