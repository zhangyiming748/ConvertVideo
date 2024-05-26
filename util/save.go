package util

import (
	"errors"
	"log"
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
		log.Printf("处理后的文件比源文件更大,放弃删除\t源文件大小:%v\t目标文件大小:%v\n", src, dst)
		return 0, errors.New("处理后的文件比源文件更大,放弃删除")
	}
	save := float64(src-dst) / MB
	return save, nil
}
