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

func ProcessVideo2VP9(in mediainfo.BasicInfo) {
	in.InsertVideoInfo()
	mi := FastMediaInfo.GetStandMediaInfo(in.FullPath)
	if strings.Contains(in.FullPath, "h265") || strings.Contains(in.FullPath, "vp9") {
		log.Printf("跳过当前已经在h265/vp9目录中的文件:%v\n", in.FullPath)
		return
	}

	FrameCount := mi.Video.FrameCount
	if mi.Video.Format == "HEVC" || mi.Video.Format == "vp09" {
		log.Println("跳过已经转码的视频")
		return
	}
	if mi.Video.CodecID == "hvc1" || mi.Video.CodecID == "vp09" {
		log.Println("跳过已经转码的视频")
		return
	}
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
	/*
		ffmpeg -threads 8 -i in.mp4 -cpu-used 8 -preset medium -c:v libvpx-vp9 -tile-columns 6 -frame-paralle 1 -crf 31 -c:a libopus -ac 1 -map_chapters -1 -threads 8 -cpu-used 8 -preset medium out.mp4"
	*/
	cmd := exec.Command("ffmpeg", "-threads", constant.GetCpuNums(), "-i", in.FullPath, "-cpu-used", "8", "-preset", "medium", "-c:v", "libvpx-vp9", "-tile-columns", "6", "-frame-parallel", "1", "-crf", crf, "-c:a", "libopus", "-vbr", "on", "-ac", "1", "-map_chapters", "-1", "-threads", constant.GetCpuNums(), "-cpu-used", "8", "-preset", "medium", mp4)
	cut := strings.Join([]string{in.PurgePath, "cut.txt"}, string(os.PathSeparator))
	if util.IsExist(cut) {
		split := util.ReadByLine(cut)
		ss := split[0]
		to := split[1]
		cmd = exec.Command("ffmpeg", "-threads", constant.GetCpuNums(), "-i", in.FullPath, "-cpu-used", "8", "-preset", "medium", "-ss", ss, "-to", to, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libopus", "-vbr", "on", "-ac", "1", "-map_chapters", "-1", "-threads", constant.GetCpuNums(), "-cpu-used", "8", "-preset", "medium", mp4)
	}
	if width > 1920 || height > 1920 {
		log.Printf("视频大于1080P需要使用其他程序先处理视频尺寸:%v\n", in)
		ResizeVideo(in)
		return
	}
	log.Printf("生成的最终命令:%v\n", cmd.String())
	msg := fmt.Sprintf("当前正在处理的视频总帧数:%v", FrameCount)
	if err := util.ExecCommand(cmd, msg); err != nil {
		return
	} else {
		log.Println("视频编码运行完成")
	}
	originsize, _ := util.GetSize(in.FullPath)
	aftersize, _ := util.GetSize(mp4)
	sub, _ := util.GetDiffSize(originsize, aftersize)
	fmt.Printf("savesize: %f MB\n", sub)
	//如果新文件比源文件还大 不删除源文件
	if aftersize < originsize {
		err := os.Remove(in.FullPath)
		if err != nil {
			log.Printf("删除失败:%v\n", in.FullPath)
			return
		} else {
			log.Printf("删除成功:%v\n", in.FullPath)
		}
	} else {
		log.Printf("转码后文件:%v\t大于源文件:%v\n", mp4, in.FullPath)
	}
	log.Printf("本次转码完成，文件大小减少 %f MB\n", sub)
}
