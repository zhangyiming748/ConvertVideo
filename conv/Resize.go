package conv

import (
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

func ResizeVideo(in mediainfo.BasicInfo) {
	vinfo := in.VInfo.Media.Track[1]
	width, _ := strconv.Atoi(vinfo.Width)
	height, _ := strconv.Atoi(vinfo.Height)

	if width > height {
		log.Printf("横屏视频:%v\n", in)
		Resize(in, "1920x1080")
	} else if width < height {
		log.Printf("竖屏视频:%v\n", in)
		Resize(in, "1080x1920")
	} else {
		log.Printf("正方形视频:%v\n", in)
		Resize(in, "1920x1920")
	}
}
func Resize(in mediainfo.BasicInfo, p string) {
	dst := in.PurgePath // 文件所在路径 不包含最后一个路径分隔符
	if strings.Contains(in.PurgePath, "resize") {
		return
	}
	dst = strings.Join([]string{dst, "resize"}, string(os.PathSeparator)) //二级目录
	fname := replace.ForFileName(in.PurgeName)                            //仅文件名
	fname = strings.Join([]string{fname, "mp4"}, ".")
	os.Mkdir(dst, 0777)
	log.Printf("输出文件夹:%v\n", dst)
	out := strings.Join([]string{dst, fname}, string(os.PathSeparator))
	log.Printf("源文件:%v\t目标文件:%v\n", in.FullPath, out)
	var cmd *exec.Cmd
	switch p {
	case "1920x1080":
		cmd = exec.Command("ffmpeg", "-i", in.FullPath, "-strict", "-2", "-vf", "scale=-1:1080", "-c:v", "libvpx-vp9", "-crf", "31", "-c:a", "libvorbis", "-ac", "1", out)
	case "1080x1920":
		cmd = exec.Command("ffmpeg", "-i", in.FullPath, "-strict", "-2", "-vf", "scale=-1:1920", "-c:v", "libvpx-vp9", "-crf", "31", "-c:a", "libvorbis", "-ac", "1", out)
	case "1920x1920":
		cmd = exec.Command("ffmpeg", "-i", in.FullPath, "-strict", "-2", "-vf", "scale=1920:1920", "-c:v", "libvpx-vp9", "-crf", "31", "-c:a", "libvorbis", "-ac", "1", out)
	default:
		log.Fatalf("不正常的视频源:%v\n", in.FullPath)
	}
	log.Printf("生成的最终命令:%v\n", cmd.String())
	frameCount := FastMediaInfo.GetStandMediaInfo(in.FullPath).Video.FrameCount
	if err := util.ExecCommand(cmd, frameCount); err != nil {
		log.Printf("resize发生错误:%v\t命令原文:%v\n", err, cmd.String())
		return
	}

	if err := os.Remove(in.FullPath); err != nil {
		log.Printf("删除失败:%n", err)
	} else {
		log.Printf("删除成功:%v\n", in.FullPath)
	}
}
