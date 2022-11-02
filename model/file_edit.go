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
	e.WriteByte([]byte(str))
}

// WriteStringLn 写入字符串内容并在之后换行
func (e *FileEdit) WriteStringLn(str string) {
	e.WriteString(str + "\n")
}

// WriteByte 写入字节内容
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
