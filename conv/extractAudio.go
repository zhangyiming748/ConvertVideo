package conv

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/GetFileInfo"
	"github.com/zhangyiming748/ConvertVideo/util"
	"log/slog"
	"os/exec"
	"strings"
)

func AllVideos2Audio(root, pattern, threads string) {
	files := GetFileInfo.GetAllFileInfo(root, pattern)
	for _, file := range files {
		Video2Audio(file, threads)
	}
}

func Videos2Audio(dir, pattern, threads string) {
	files := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, file := range files {
		Video2Audio(file, threads)
	}
}
func Video2Audio(in GetFileInfo.BasicInfo, threads string) {
	out := strings.Replace(in.FullPath, in.PurgeExt, "ogg", 1)
	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-vn", "-ac", "1", out)
	slog.Info("生成的命令", slog.String("command", fmt.Sprint(cmd)))
	if err := util.ExecCommand(cmd); err != nil {
		slog.Warn("命令执行中出现错误")
	}
	slog.Debug("视频提取音频运行完成")
	//if err == nil {
	//	if err = os.RemoveAll(in.FullPath); err != nil {
	//		slog.Warn("删除失败", slog.Any("源文件", in.FullPath), slog.Any("错误", err))
	//	} else {
	//		slog.Info("删除成功", slog.Any("源文件", in.FullName))
	//	}
	//}
}
