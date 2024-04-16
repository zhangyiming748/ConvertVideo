package util

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/constant"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

/*
执行命令过程中可以循环打印消息
*/
func ExecCommand(c *exec.Cmd, msg string) (e error) {
	slog.Info("开始执行命令", slog.String("命令原文", fmt.Sprint(c)))
	if level := constant.GetLevel(); level == "Debug" {
		stdout, err := c.StdoutPipe()
		c.Stderr = c.Stdout
		if err != nil {
			slog.Warn("连接Stdout产生错误", slog.String("命令原文", fmt.Sprint(c)), slog.String("错误原文", fmt.Sprint(err)))
			return err
		}
		if err = c.Start(); err != nil {
			slog.Warn("启动cmd命令产生错误", slog.String("命令原文", fmt.Sprint(c)), slog.String("错误原文", fmt.Sprint(err)))
			return err
		}
		for {
			tmp := make([]byte, 1024)
			_, err := stdout.Read(tmp)
			t := string(tmp)
			t = strings.Replace(t, "\u0000", "", -1)
			fmt.Println(t)
			fmt.Println(msg)
			if err != nil {
				break
			}
		}
		if err = c.Wait(); err != nil {
			slog.Warn("命令执行中产生错误", slog.String("命令原文", fmt.Sprint(c)), slog.String("错误原文", fmt.Sprint(err)))
			return err
		}
	} else {
		fmt.Println(msg)
		if output, err := c.CombinedOutput(); err != nil {
			slog.Warn("命令执行中产生错误", slog.String("命令原文", fmt.Sprint(c)), slog.String("错误原文", fmt.Sprint(err)))
			return err
		} else {
			slog.Debug("命令执行完毕", slog.String("输出", string(output)))
		}
	}
	if isExitLabel() {
		slog.Debug("命令端获取到退出状态,命令结束后退出", slog.String("最后一条命令", fmt.Sprint(c)))
		os.Exit(0)
	}
	return nil
}

/*
判断古希腊掌管退出信号的文件是否存在
*/
func isExitLabel() bool {
	filePath := "/exit"

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("古希腊掌管退出信号的文件不存在")
		return false
	} else {
		fmt.Println("古希腊掌管退出信号的文件存在")
		return true
	}
}
