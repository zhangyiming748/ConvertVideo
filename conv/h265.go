package conv

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/replace"
	"github.com/zhangyiming748/ConvertVideo/util"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ProcessVideo2H265(in mediainfo.BasicInfo) {
	in.InsertVideoInfo()
	var (
		width  int
		height int
	)
	if strings.Contains(in.FullPath, "h265") || strings.Contains(in.FullPath, "vp9") {
		log.Printf("跳过当前已经在h265/vp9目录中的文件:%v\n", in.FullPath)
		return
	}
	FrameCount := ""
	for _, v := range in.VInfo.Media.Track {
		if v.Type == "Video" {
			vinfo := v
			log.Printf("编码:%v\tCodecID:%v\n", vinfo.Format, vinfo.CodecID)
			if vinfo.Format == "HEVC" || vinfo.CodecID == "hvc1" || vinfo.Format == "vp09" || vinfo.CodecID == "vp09" {
				log.Println("跳过已经转码的视频")
				return
			}
			width, _ = strconv.Atoi(vinfo.Width)
			height, _ = strconv.Atoi(vinfo.Height)
			log.Printf("获取帧数:%v\n", vinfo.FrameCount)
			FrameCount = vinfo.FrameCount
		}
	}
	defer func() {
		if err := recover(); err != nil {
			log.Printf("处理视频失败:%v\n", err)
		} else {
			log.Printf("处理视频成功:%v\n", in.FullPath)
			if err = os.Remove(in.FullPath); err != nil {
				log.Printf("删除失败\t源文件:%v\t错误:%v\n ", in.FullPath, err)
			} else {
				log.Printf("删除成功\t源文件:%v\n", in.FullName)
			}
		}
	}()
	middle := "h265"
	if err := os.Mkdir(strings.Join([]string{in.PurgePath, middle}, string(os.PathSeparator)), 0777); err != nil {
		if strings.Contains(err.Error(), "file exists") {
			log.Println("输出文件夹已存在")
		}
	} else {
		log.Println("创建输出文件夹")
	}
	dstPurgeName := replace.ForFileName(in.PurgeName) // 输入文件格式化后的新文件名
	out := strings.Join([]string{in.PurgePath, string(os.PathSeparator), middle, string(os.PathSeparator), dstPurgeName, ".mp4"}, "")
	defer func() {
		if err := recover(); err != nil {
			log.Printf("出现错误\t输入文件:%v\t输出文件%v\n:", in.FullPath, out)
		}
	}()
	mp4 := strings.Replace(out, in.PurgeExt, "mp4", -1)
	log.Printf("输入文件:%v\t输出文件%v\n:", in.FullPath, out)
	cmd := exec.Command("ffmpeg", "-i", in.FullPath, "-c:v", "libx265", "-crf", "22", "-tag:v", "hvc1", "-c:a", "libvorbis", "-ac", "1", "-map_chapters", "-1", mp4)
	if width > 1920 && height > 1920 {
		log.Printf("视频大于1080P需要使用其他程序先处理视频尺寸:%v\n", in)
		ResizeVideo(in)
		return
	}
	log.Printf("生成的命令", cmd.String())
	msg := fmt.Sprintf("当前正在处理的视频总帧数:%v", FrameCount)
	util.ExecCommand(cmd, msg)
	log.Println("视频编码运行完成")
	originsize, _ := util.GetSize(in.FullPath)
	aftersize, _ := util.GetSize(mp4)
	sub, _ := util.GetDiffSize(originsize, aftersize)
	fmt.Printf("savesize: %f MB\n", sub)
	log.Printf("本次转码完成，文件大小减少 %f MB\n", sub)
}
