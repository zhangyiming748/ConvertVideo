package mediainfo

import (
	"path/filepath"
	"strings"
)

type BasicInfo struct {
	FullPath  string    //文件的绝对路径
	PurgePath string    //文件的路径部分
	FullName  string    //完整文件名
	PurgeName string    //文件名的部分
	PurgeExt  string    //扩展名的部分
	VInfo     VideoInfo //视频文件的详细信息
}

func GetBasicInfo(fp string) *BasicInfo {
	binfo := new(BasicInfo)
	binfo.FullPath = fp
	dir, file := filepath.Split(fp)
	binfo.PurgePath = strings.TrimSuffix(dir, string(filepath.Separator))
	binfo.FullName = file
	binfo.PurgeExt = strings.Replace(filepath.Ext(file), ".", "", 1)
	binfo.PurgeName = strings.TrimSuffix(file, binfo.PurgeExt)
	return binfo
}
