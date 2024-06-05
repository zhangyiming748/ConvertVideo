package conv

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/constant"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/replace"
	"github.com/zhangyiming748/ConvertVideo/util"
	"github.com/zhangyiming748/FastMediaInfo"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ProcessVideo2clip(in mediainfo.BasicInfo) {
	in.InsertVideoInfo()
	mi := FastMediaInfo.GetStandMediaInfo(in.FullPath)
	if strings.Contains(in.FullPath, "h265") || strings.Contains(in.FullPath, "vp9") {
		log.Printf("跳过当前已经在h265/vp9目录中的文件:%v\n", in.FullPath)
		return
	}

	FrameCount := mi.Video.FrameCount

	log.Printf("fullname:%v\n", in.FullName)
	middle := "vp9"
	if exist := os.Mkdir(strings.Join([]string{in.PurgePath, middle}, string(os.PathSeparator)), 0777); exist != nil {
		if strings.Contains(exist.Error(), "file exists") {
			log.Println("输出文件夹已存在")
		}
	} else {
		log.Println("创建输出文件夹")
	}
	dstPurgeName := replace.ForFileName(in.PurgeName) // 输入文件格式化后的新文件名
	out := strings.Join([]string{in.PurgePath, string(os.PathSeparator), middle, string(os.PathSeparator), dstPurgeName, ".mp4"}, "")
	mp4 := strings.Replace(out, in.PurgeExt, "mp4", -1)
	log.Printf("输入文件:%v\t输出文件:%v\n", in.FullPath, mp4)
	var (
		width, _  = strconv.Atoi(mi.Video.Width)
		height, _ = strconv.Atoi(mi.Video.Height)
	)
	crf := FastMediaInfo.GetCRF("vp9", width, height)
	if crf == "" {
		crf = "31"
		log.Printf("没有查询到crf,使用默认crf:%v\n", crf)
	}
	cut := strings.Join([]string{in.PurgePath, "cut.txt"}, string(os.PathSeparator))
	if util.IsExist(cut) {
		split := util.ReadByLine(cut)
		lenght := len(split)
		count := 0
		for i := 0; i < lenght-1; i += 2 {
			count++
			part := strings.Join([]string{"part", strconv.Itoa(count), ".mp4"}, "")
			clip := strings.Replace(mp4, ".mp4", part, 1)
			ss := split[i]
			to := split[i+1]
			cmd := exec.Command("ffmpeg", "-threads", constant.GetCpuNums(), "-i", in.FullPath, "-cpu-used", "8", "-ss", ss, "-to", to, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libopus", "-vbr", "on", "-ac", "1", "-map_chapters", "-1", "-threads", constant.GetCpuNums(), clip)
			log.Printf("当前生成的命令:%v\n", cmd.String())
			msg := fmt.Sprintf("当前正在处理的视频总帧数:%v", FrameCount)
			if err := util.ExecCommand(cmd, msg); err != nil {
				return
			} else {
				log.Println("视频编码运行完成")
			}
		}
	}
}
