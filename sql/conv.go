package sql

import (
	"gorm.io/gorm"
)

type Conv struct {
	gorm.Model
	OriginName string `gorm:"origin;comment:'初始文件名'"`
	AfterName  string `gorm:"after;comment:'新文件名'"`
	Src        string `gorm:"src;comment:'源文件'"`
	Dst        string `gorm:"dst;comment:'目标文件'"`
	SrcSize    string `gorm:"srcsize;comment:'源文件大小'"`
	DstSize    string `gorm:"dstsize;comment:'目标文件大小'"`
	Srt        string `gorm:"type:text;comment:'非实时命令输出的原文'"`
}

func (c *Conv) FindOneByOriginName() *gorm.DB {
	return GetEngine().Where("name = ?", c.OriginName).First(&c)
}

func (c *Conv) SetOne() *gorm.DB {
	return GetEngine().Create(&c)
}
