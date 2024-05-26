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

func RotateVideo(in mediainfo.BasicInfo, direction string) {

	if strings.Contains(in.PurgePath, "rotate") {
		return
	}
	in.InsertVideoInfo()
	log.Printf("插入细节信息后:%v\n", in)
	dst := strings.Join([]string{in.PurgePath, "rotate"}, string(os.PathSeparator))
	os.Mkdir(dst, os.ModePerm)
	FrameCount := ""
	for _, v := range in.VInfo.Media.Track {
		if v.Type == "Video" {
			FrameCount = v.FrameCount
		}
	}
	fname := in.PurgeName
	fname = replace.ForFileName(fname)
	fname = strings.Join([]string{fname, "mp4"}, ".")
	out := strings.Join([]string{dst, fname}, string(os.PathSeparator))
	var cmd *exec.Cmd
	var transport string
	switch direction {
	case "ToRight":
		transport = "transpose=1"
	case "ToLeft":
		transport = "transpose=2"
	default:
		return
	}
	var (
		width  int
		height int
	)
	for _, v := range in.VInfo.Media.Track {
		if v.Type == "Video" {
			width, _ = strconv.Atoi(v.Width)
			height, _ = strconv.Atoi(v.Height)
			log.Printf("分辨率:%v x %v\n", width, height)
		}
	}
	crf := util.GetCrf("h265", width, height)
	if crf == "" {
		crf = "31"
		log.Printf("没有查询到crf,使用默认crf:%v\n", crf)
	}
	log.Printf("获取到的crf=%v\n", crf)
	cmd = exec.Command("ffmpeg", "-y", "-i", in.FullPath, "-vf", transport, "-c:v", "libx265", "-crf", crf, "-c:a", "libvorbis", "-ac", "1", "-map_chapters", "-1", out)
	err := util.ExecCommand(cmd, FrameCount)
	if err != nil {
		os.Exit(-1)
	}
	originsize, _ := util.GetSize(in.FullPath)
	aftersize, _ := util.GetSize(out)
	sub, _ := util.GetDiffSize(originsize, aftersize)
	fmt.Printf("savesize: %f MB\n", sub)
	if aftersize < originsize {
		if err = os.Remove(in.FullPath); err != nil {
			log.Printf("删除失败:%v\n", in.FullPath)
		} else {
			log.Printf("删除成功:%v\n", in.FullPath)
		}
	} else {
		log.Printf("转码后文件:%v\t大于源文件:%v\n,保留不删除", out, in.FullPath)
	}
	log.Printf("本次转码完成，文件大小减少 %f MB\n", sub)

}
