package conv

import (
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/util"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

/*
mkv添加ass
*/
func MkvWithAss(in mediainfo.BasicInfo) {
	srt := strings.Replace(in.FullPath, in.PurgeExt, "srt", 1)
	//ass := strings.Replace(file.FullPath, ".mp4", ".ass", 1)
	if util.IsExist(srt) {
		// output := strings.Replace(file, ".mp4", "_with_subtitle.mp4", 1)
		ext := path.Ext(in.FullPath)
		output := strings.Replace(in.FullPath, ext, "_with_subtitle.mkv", 1)
		//cmd := exec.Command("ffmpeg", "-i", file.FullPath, "-f", "srt", "-i", srt, "-c:v", "libx265", "-c:a", "aac", "-ac", "1", "-tag:v", "hvc1", "-c:s", "mov_text", output)
		// ffmpeg -i input.mkv -i input.ass -c copy -c:s ass output.mkv
		cmd := exec.Command("ffmpeg", "-i", in.FullPath, "-i", srt, "-c:v", "libvpx-vp9", "-c:a", "libvorbis", "-ac", "1", "-c:s", "ass", output)
		log.Printf("生成的命令: %s\n", cmd.String())
		combinedOutput, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("命令执行失败: %s\n", err.Error())
			return
		} else {
			log.Printf("命令成功执行: %s\n", string(combinedOutput))
			os.Remove(in.FullPath)
		}
	}
}
