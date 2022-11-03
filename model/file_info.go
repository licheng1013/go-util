package model

// FileInfo 文件信息封装
type FileInfo struct {
	// 文件名
	FileName string `from:"fileName" json:"fileName"`
	// 1 是目录，0 默认文件
	IsDirectory int8 `from:"isDirectory" json:"isDirectory"`
	// 文件绝对路径
	FilePath string `json:"filePath"`
	// 修改时间
	UpdatedAt string `json:"updatedAt"`
}
