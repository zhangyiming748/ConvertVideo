package util

import (
	"errors"
	"log/slog"
	"os"
)

const (
	MB = 1024 * 1024
)

/*
计算所提供文件大小 字节
*/

func GetSize(fp string) (int64, error) {
	if file, err := os.Open(fp); err != nil {
		return 0, err
	} else if info, err := file.Stat(); err != nil {
		return 0, err
	} else {
		defer file.Close()
		size := info.Size()
		return size, nil
	}
}

/*
计算所给定的两个文件大小差 返回MB
*/
func GetDiffSize(src, dst int64) (float64, error) {
	if dst >= src {
		slog.Warn("处理后的文件比源文件更大,放弃删除", slog.Int64("源文件大小", src), slog.Int64("目标文件大小", dst))
		return 0, errors.New("处理后的文件比源文件更大,放弃删除")
	}
	save := float64(src-dst) / MB
	return save, nil
}
