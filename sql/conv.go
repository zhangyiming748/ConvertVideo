package sql

import (
	"gorm.io/gorm"
)

type Conv struct {
	gorm.Model
	Src      string `gorm:"src;comment:'源文件名'"`
	Dst      string `gorm:"dst;comment:'目标文件名'"`
	SrcSize  int64  `gorm:"srcsize;comment:'源文件大小';type:int64"`
	DstSize  int64  `gorm:"dstsize;comment:'目标文件大小';type:int64"`
	IsBigger bool   `gorm:"dstsize;comment:'新文件是否比旧文件还大';type:bool"`

	//Srt        string `gorm:"type:text;comment:'非实时命令输出的原文'"`
}

func (c *Conv) FindOneByOriginName() *gorm.DB {
	return GetEngine().Where("name = ?", c.Src).First(&c)
}

func (c *Conv) SetOne() *gorm.DB {
	return GetEngine().Create(&c)
}
