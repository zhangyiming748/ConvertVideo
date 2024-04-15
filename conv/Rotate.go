package conv

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/replace"
	"github.com/zhangyiming748/ConvertVideo/util"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func RotateVideo(in mediainfo.BasicInfo, direction string) {
	if strings.Contains(in.PurgePath, "rotate") {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			if err = os.Remove(in.FullPath); err != nil {
				slog.Warn("删除失败", slog.Any("源文件", in.FullPath), slog.Any("错误", err))
			} else {
				slog.Debug("删除成功", slog.Any("源文件", in.FullPath))
			}
		}
	}()
	dst := strings.Join([]string{in.PurgePath, "rotate"}, string(os.PathSeparator))
	os.Mkdir(dst, os.ModePerm)
	FrameCount := ""
	for _, v := range in.VInfo.Media.Track {
		if v.Type == "video" {
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
		if v.Type == "video" {
			width, _ = strconv.Atoi(v.Width)
			height, _ = strconv.Atoi(v.Height)
		}
	}
	crf := util.GetCrf("vp9", width, height)
	cmd = exec.Command("ffmpeg", "-i", in.FullPath, "-vf", transport, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libvorbis", "-ac", "1", "-map_chapters", "-1", out)
	util.ExecCommand(cmd, FrameCount)
	originsize, _ := util.GetSize(in.FullPath)
	aftersize, _ := util.GetSize(out)
	sub, _ := util.GetDiffSize(originsize, aftersize)
	fmt.Printf("savesize: %f MB\n", sub)

	slog.Info(fmt.Sprintf("本次转码完成，文件大小减少 %f MB\n", sub))

}
