package conv

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/constant"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/util"
	"github.com/zhangyiming748/FastMediaInfo"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

/*
mkv添加ass
*/
func MkvWithAss(in mediainfo.BasicInfo) {
	mi := FastMediaInfo.GetStandMediaInfo(in.FullPath)
	FrameCount := mi.Video.FrameCount
	var (
		width, _  = strconv.Atoi(mi.Video.Width)
		height, _ = strconv.Atoi(mi.Video.Height)
	)
	crf := FastMediaInfo.GetCRF("vp9", width, height)
	if crf == "" {
		crf = "31"
		log.Printf("没有查询到crf,使用默认crf:%v\n", crf)
	}
	srt := strings.Replace(in.FullPath, in.PurgeExt, "srt", 1)
	if util.IsExist(srt) {
		ext := path.Ext(in.FullPath)
		output := strings.Replace(in.FullPath, ext, "_with_subtitle.mkv", 1)
		var cmd *exec.Cmd
		cmd = exec.Command("ffmpeg", "-threads", constant.GetCpuNums(), "-i", in.FullPath, "-i", srt, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libopus", "-ac", "1", "-c:s", "ass", "-threads", constant.GetCpuNums(), output)
		log.Printf("生成的命令: %s\n", cmd.String())
		msg := fmt.Sprintf("当前正在处理的视频总帧数:%v", FrameCount)
		err := util.ExecCommand(cmd, msg)
		if err != nil {
			log.Printf("命令执行失败: %s\n", err.Error())
			return
		} else {
			log.Printf("命令成功执行: %s\n", cmd.String())
			os.Remove(in.FullPath)
		}
	}
}
