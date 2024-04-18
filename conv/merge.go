package conv

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

/*
mkv添加ass
*/
func MkvWithAss(in mediainfo.BasicInfo) {
	srt := strings.Replace(in.FullPath, in.PurgeExt, "srt", 1)
	//ass := strings.Replace(file.FullPath, ".mp4", ".ass", 1)
	if isExist(srt) {
		// output := strings.Replace(file, ".mp4", "_with_subtitle.mp4", 1)
		ext := path.Ext(in.FullPath)
		output := strings.Replace(in.FullPath, ext, "_with_subtitle.mkv", 1)
		//cmd := exec.Command("ffmpeg", "-i", file.FullPath, "-f", "srt", "-i", srt, "-c:v", "libx265", "-c:a", "aac", "-ac", "1", "-tag:v", "hvc1", "-c:s", "mov_text", output)
		// ffmpeg -i input.mkv -i input.ass -c copy -c:s ass output.mkv
		cmd := exec.Command("ffmpeg", "-i", in.FullPath, "-i", srt, "-c:v", "libvpx-vp9", "-c:a", "libvorbis", "-ac", "1", "-c:s", "ass", output)
		fmt.Printf("生成的命令: %s\n", cmd.String())
		combinedOutput, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("命令执行失败: %s\n", err.Error())
			return
		} else {
			fmt.Printf("命令成功执行: %s\n", string(combinedOutput))
			os.Remove(in.FullPath)
		}
	}
}

func isExist(fp string) bool {
	_, err := os.Stat(fp)
	if os.IsNotExist(err) {
		fmt.Printf("%s 对应的字幕文件不存在\n", fp)
		return false
	} else {
		fmt.Printf("%s 对应的字幕文件存在\n", fp)
		return true
	}
}
func getFilesWithExtension(folderPath string, extension string) ([]string, error) {
	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
