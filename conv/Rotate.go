package conv

import (
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
	_ = os.Mkdir(dst, os.ModePerm)
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
	cmd = exec.Command("ffmpeg", "-y", "-i", in.FullPath, "-vf", transport, "-c:v", "libx265", "-c:a", "libopus", "-ac", "1", "-map_chapters", "-1", out)
	err := util.ExecCommand(cmd, FrameCount)
	if err != nil {
		log.Fatalln("旋转视频命令发生错误")
	}
	originsize, _ := util.GetSize(in.FullPath)
	aftersize, _ := util.GetSize(out)
	sub, _ := util.GetDiffSize(originsize, aftersize)
	log.Printf("转换前%fM转换后%fM节省%fM\n", originsize/util.MB, aftersize/util.MB, sub)
	if err = os.Remove(in.FullPath); err != nil {
		log.Printf("删除失败:%v\n", in.FullPath)
	} else {
		log.Printf("删除成功:%v\n", in.FullPath)
	}
}
