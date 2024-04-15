package conv

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/replace"
	"github.com/zhangyiming748/ConvertVideo/sql"
	"github.com/zhangyiming748/ConvertVideo/util"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ProcessVideo2VP9(in mediainfo.BasicInfo) {
	in.InsertVideoInfo()
	c := new(sql.Conv)
	defer c.SetOne()
	var (
		width  int
		height int
	)
	if strings.Contains(in.FullPath, "h265") || strings.Contains(in.FullPath, "vp9") {
		slog.Debug("跳过当前已经在h265/vp9目录中的文件", slog.String("文件名", in.FullPath))
		return
	}
	safeDelete := false
	FrameCount := ""
	for _, v := range in.VInfo.Media.Track {
		if v.Type == "Video" {
			vinfo := v
			slog.Info("编码", slog.String("Format", vinfo.Format), slog.String("CodecID", vinfo.CodecID))
			if vinfo.Format == "HEVC" || vinfo.CodecID == "hvc1" || vinfo.Format == "vp09" || vinfo.CodecID == "vp09" {
				slog.Info("跳过已经转码的视频")
				return
			}
			width, _ = strconv.Atoi(vinfo.Width)
			height, _ = strconv.Atoi(vinfo.Height)
			slog.Info("获取帧数", slog.String("当前视频帧数", vinfo.FrameCount))
			FrameCount = vinfo.FrameCount
		}
	}
	defer func() {
		if err := recover(); err != nil {
			slog.Error("处理视频失败", slog.Any("错误", err))
		} else {
			//如果可以安全删除
			if safeDelete {
				slog.Info("处理视频成功", slog.String("文件名", in.FullPath))
				if err = os.Remove(in.FullPath); err != nil {
					slog.Warn("删除失败", slog.Any("源文件", in.FullPath), slog.Any("错误", err))
				} else {
					slog.Debug("删除成功", slog.Any("源文件", in.FullName))
				}
			}
		}
	}()

	//slog.Debug("文件信息", slog.Any("info", in))

	slog.Debug("fullname", slog.String("fullname", in.FullName))
	middle := "vp9"
	if err := os.Mkdir(strings.Join([]string{in.PurgePath, middle}, string(os.PathSeparator)), 0777); err != nil {
		if strings.Contains(err.Error(), "file exists") {
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
	crf := util.GetCrf("vp9", width, height)
	cmd := exec.Command("ffmpeg", "-i", in.FullPath, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libvorbis", "-ac", "1", "-map_chapters", "-1", mp4)
	if width > 1920 && height > 1920 {
		slog.Warn("视频大于1080P需要使用其他程序先处理视频尺寸", slog.Any("原视频", in))
		ResizeVideo(in)
		return
	}
	slog.Info("生成的命令", slog.String("command", fmt.Sprint(cmd)))
	msg := fmt.Sprintf("当前正在处理的视频总帧数:%v", FrameCount)
	util.ExecCommand(cmd, msg)
	slog.Debug("视频编码运行完成")

	originsize, _ := util.GetSize(in.FullPath)
	aftersize, _ := util.GetSize(mp4)
	sub, _ := util.GetDiffSize(originsize, aftersize)
	fmt.Printf("savesize: %f MB\n", sub)
	//todo 如果新文件比源文件还大 不删除源文件
	if aftersize < originsize {
		safeDelete = true
	}

	c.Src = in.FullPath
	c.Dst = mp4
	c.SrcSize = originsize
	c.DstSize = aftersize
	if !safeDelete {
		c.IsBigger = true
	} else {
		c.IsBigger = false
	}

	slog.Info(fmt.Sprintf("本次转码完成，文件大小减少 %f MB\n", sub))

}
