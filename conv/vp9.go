package conv

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/constant"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/replace"
	"github.com/zhangyiming748/ConvertVideo/util"
	"github.com/zhangyiming748/FastMediaInfo"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ProcessVideo2VP9(in mediainfo.BasicInfo) {
	in.InsertVideoInfo()
	mi := FastMediaInfo.GetStandMediaInfo(in.FullPath)
	if strings.Contains(in.FullPath, "h265") || strings.Contains(in.FullPath, "vp9") {
		slog.Debug("跳过当前已经在h265/vp9目录中的文件", slog.String("文件名", in.FullPath))
		return
	}
	FrameCount := mi.Video.FrameCount
	if mi.Video.Format == "HEVC" || mi.Video.Format == "vp09" {
		slog.Info("跳过已经转码的视频")
		return
	}
	if mi.Video.CodecID == "hvc1" || mi.Video.CodecID == "vp09" {
		slog.Info("跳过已经转码的视频")
		return
	}
	slog.Debug("fullname", slog.String("fullname", in.FullName))
	middle := "vp9"
	if exist := os.Mkdir(strings.Join([]string{in.PurgePath, middle}, string(os.PathSeparator)), 0777); exist != nil {
		if strings.Contains(exist.Error(), "file exists") {
			slog.Debug("输出文件夹已存在")
		}
	} else {
		slog.Debug("创建输出文件夹")
	}
	dstPurgeName := replace.ForFileName(in.PurgeName) // 输入文件格式化后的新文件名
	out := strings.Join([]string{in.PurgePath, string(os.PathSeparator), middle, string(os.PathSeparator), dstPurgeName, ".mp4"}, "")
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("出现错误", slog.String("输入文件", in.FullPath), slog.String("输出文件", out))
		}
	}()
	slog.Debug("", slog.String("out", out), slog.String("extName", in.PurgeExt))
	mp4 := strings.Replace(out, in.PurgeExt, "mp4", -1)
	slog.Debug("调试", slog.String("输入文件", in.FullPath), slog.String("输出文件", mp4))
	var (
		width, _  = strconv.Atoi(mi.Video.Width)
		height, _ = strconv.Atoi(mi.Video.Height)
	)
	crf := FastMediaInfo.GetCRF("vp9", width, height)
	if crf == "" {
		crf = "31"
		slog.Warn("没有查询到crf", slog.String("使用默认crf", crf))
	}
	cmd := exec.Command("ffmpeg", "-threads", constant.GetCpuNums(), "-i", in.FullPath, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libvorbis", "-ac", "1", "-map_chapters", "-1", "-threads", constant.GetCpuNums(), mp4)
	if width > 1920 || height > 1920 {
		slog.Warn("视频大于1080P需要使用其他程序先处理视频尺寸", slog.Any("原视频", in))
		ResizeVideo(in)
		return
	}
	slog.Info("生成的命令", slog.String("command", fmt.Sprint(cmd)))
	msg := fmt.Sprintf("当前正在处理的视频总帧数:%v", FrameCount)
	if err := util.ExecCommand(cmd, msg); err != nil {
		return
	} else {
		slog.Info("命令成功执行,删除源文件", slog.String("command", fmt.Sprint(cmd)))
	}
	slog.Debug("视频编码运行完成")
	originsize, _ := util.GetSize(in.FullPath)
	aftersize, _ := util.GetSize(mp4)
	sub, _ := util.GetDiffSize(originsize, aftersize)
	fmt.Printf("savesize: %f MB\n", sub)
	//如果新文件比源文件还大 不删除源文件
	if aftersize < originsize {
		err := os.Remove(in.FullPath)
		if err != nil {
			slog.Warn("删除失败", slog.String("文件", in.FullPath))
			return
		} else {
			slog.Info("删除成功", slog.String("文件", in.FullPath))
		}
	}
	slog.Info(fmt.Sprintf("本次转码完成，文件大小减少 %f MB\n", sub))
}
