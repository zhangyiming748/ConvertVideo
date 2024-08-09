package conv

import (
	"github.com/zhangyiming748/ConvertVideo/constant"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/replace"
	"github.com/zhangyiming748/ConvertVideo/util"
	"github.com/zhangyiming748/DeepLX"
	"github.com/zhangyiming748/FastMediaInfo"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ProcessVideo2H265(in mediainfo.BasicInfo) {
	in.InsertVideoInfo()
	var (
		width  int
		height int
	)
	mi := FastMediaInfo.GetStandMediaInfo(in.FullPath)
	FrameCount := mi.Video.FrameCount
	width, _ = strconv.Atoi(mi.Video.Width)
	height, _ = strconv.Atoi(mi.Video.Height)
	if mi.Video.Format == "HEVC" || mi.Video.Format == "vp09" {
		log.Println("跳过已经转码的视频")
		return
	}
	if mi.Video.CodecID == "hvc1" || mi.Video.CodecID == "vp09" {
		log.Println("跳过已经转码的视频")
		return
	}
	middle := "h265"
	if err := os.Mkdir(strings.Join([]string{in.PurgePath, middle}, string(os.PathSeparator)), 0777); err != nil {
		if strings.Contains(err.Error(), "file exists") {
			log.Println("输出文件夹已存在")
		}
	} else {
		log.Println("创建输出文件夹")
	}
	dstPurgeName := replace.ForFileName(in.PurgeName) // 输入文件格式化后的新文件名
	if constant.GetTransTitle() {
		if dst, err := DeepLx.TranslateByDeepLX("auto", "zh", dstPurgeName, ""); err == nil {
			dstPurgeName = dst
		}
	}
	out := strings.Join([]string{in.PurgePath, string(os.PathSeparator), middle, string(os.PathSeparator), dstPurgeName, ".mp4"}, "")
	mp4 := strings.Replace(out, in.PurgeExt, "mp4", -1)

	cmd := exec.Command("ffmpeg", "-i", in.FullPath, "-c:v", "libx265", "-tag:v", "hvc1", "-c:a", "libopus", "-ac", "1", "-map_chapters", "-1", mp4)
	if runtime.GOOS == "linux" && runtime.GOARCH == "arm64" {
		cmd = exec.Command("ffmpeg", "-i", in.FullPath, "-threads", "1", "-c:v", "libx265", "-tag:v", "hvc1", "-c:a", "libopus", "-ac", "1", "-map_chapters", "-1", "-threads", "1", mp4)
	}
	if width > 1920 && height > 1920 {
		log.Printf("视频大于1080P需要使用其他程序先处理视频尺寸:%v\n", in)
		ResizeVideo(in)
		return
	}
	log.Printf("生成的命令:%v\n", cmd.String())
	if err := util.ExecCommand(cmd, FrameCount); err != nil {
		return
	}
	log.Println("视频编码运行完成")
	originsize, _ := util.GetSize(in.FullPath)
	aftersize, _ := util.GetSize(mp4)
	sub, _ := util.GetDiffSize(originsize, aftersize)
	log.Printf("转换前%fM转换后%fM节省%fM\n", originsize/util.MB, aftersize/util.MB, sub)
	if err := os.Remove(in.FullPath); err != nil {
		log.Printf("删除失败:%v\n", in.FullPath)
	} else {
		log.Printf("删除成功:%v\n", in.FullPath)
	}
}
