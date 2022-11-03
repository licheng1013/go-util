package model

import "time"

// FileInfo 文件信息封装
type FileInfo struct {
	// 文件名
	FileName string `from:"fileName" json:"fileName"`
	// 1 是目录，0 默认文件
	IsDirectory int8 `from:"isDirectory" json:"isDirectory"`
	// 文件绝对路径
	FilePath string `json:"filePath"`
	// 修改时间
	UpdatedAt time.Time
}

// File 用于扩展内部的实现 TODO 这块功能还没适配
type File interface {
	GetFileName() string
	IsDirectory() int8
	GetFilePath() string
	GetUpdatedAt() time.Time
}
