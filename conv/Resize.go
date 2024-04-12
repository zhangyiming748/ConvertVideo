package conv

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/mediaInfo"
	"github.com/zhangyiming748/ConvertVideo/replace"
	"github.com/zhangyiming748/ConvertVideo/util"
	"log/slog"
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
		slog.Debug("横屏视频", slog.Any("视频信息", in))
		Resize(in, "1920x1080")
	} else if width < height {
		slog.Debug("竖屏视频", slog.Any("视频信息", in))
		Resize(in, "1080x1920")
	} else {
		slog.Debug("正方形视频", slog.Any("视频信息", in))
		Resize(in, "1920x1920")
	}
}
func Resize(in mediainfo.BasicInfo, p string) {
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("错误", slog.String("文件信息", in.FullPath))
		}
	}()

	dst := in.PurgePath // 文件所在路径 不包含最后一个路径分隔符
	if strings.Contains(in.PurgePath, "resize") {
		return
	}
	dst = strings.Join([]string{dst, "resize"}, string(os.PathSeparator)) //二级目录
	fname := replace.ForFileName(in.PurgeName)                            //仅文件名
	fname = strings.Join([]string{fname, "mp4"}, ".")
	os.Mkdir(dst, 0777)
	slog.Debug("新建文件夹", slog.String("全名", dst))
	out := strings.Join([]string{dst, fname}, string(os.PathSeparator))
	slog.Debug("io", slog.String("源文件:", in.FullPath), slog.String("输出文件:", out))
	var cmd *exec.Cmd
	switch p {
	case "1920x1080":
		cmd = exec.Command("ffmpeg", "-i", in.FullPath, "-strict", "-2", "-vf", "scale=-1:1080", "-c:v", "libvpx-vp9", "-crf", "31", "-c:a", "libvorbis", "-ac", "1", out)
	case "1080x1920":
		cmd = exec.Command("ffmpeg", "-i", in.FullPath, "-strict", "-2", "-vf", "scale=-1:1920", "-c:v", "libvpx-vp9", "-crf", "31", "-c:a", "libvorbis", "-ac", "1", out)
	case "1920x1920":
		cmd = exec.Command("ffmpeg", "-i", in.FullPath, "-strict", "-2", "-vf", "scale=1920:1920", "-c:v", "libvpx-vp9", "-crf", "31", "-c:a", "libvorbis", "-ac", "1", out)
	default:
		slog.Warn("不正常的视频源", slog.Any("视频信息", in.FullPath))
	}
	slog.Debug("ffmpeg", slog.String("生成的命令", fmt.Sprintf("生成的命令是:%s", cmd)))
	if err := util.ExecCommand(cmd); err != nil {
		slog.Warn("resize发生错误", slog.String("命令原文", fmt.Sprint(cmd)), slog.String("错误原文", fmt.Sprint(err)), slog.String("源文件", in.FullPath))
		return
	}

	if err := os.Remove(in.FullPath); err != nil {
		slog.Warn("删除失败", slog.String("源文件", in.FullPath), slog.Any("错误文本", err))
	} else {
		slog.Warn("删除成功", slog.String("源文件", in.FullPath))
	}
}
