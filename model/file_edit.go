package model

import (
	"log"
	"os"
)

// FileEdit 文件编辑
type FileEdit struct {
	// 文件
	File *os.File
	// 文件路径
	Path string
	// 文件名
	Name string
}

// WriteString 写入字符串内容
func (e *FileEdit) WriteString(str string) {
	_, err := e.File.WriteString(str)
	if err != nil {
		panic(err)
	}
}

// WriteByte 写入字符串内容
func (e *FileEdit) WriteByte(by []byte) {
	_, err := e.File.Write(by)
	if err != nil {
		panic(err)
	}
}

// Close 关闭文件流
func (e *FileEdit) Close() {
	err := e.File.Close()
	if err != nil {
		log.Println(err.Error())
	}
}
