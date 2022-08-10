/*
 *                                  Apache License
 *                            Version 2.0, January 2004
 *                         http://www.apache.org/licenses/
 */

package util

import (
	"log"
	"os"
	"strings"
)

var pathSeparator = string(os.PathSeparator)

func ReadByte(file *os.File) []byte {
	var data []byte
	var bytes = make([]byte, 1024*8)
	for {
		n, err := file.Read(bytes)
		if err != nil {
			panic(err)
		}
		if n == 0 {
			break
		}
		// 将读取到的结果追加到data切片中
		data = append(data, bytes[:n]...)
	}
	return bytes
}

// Mkdir 创建目录,会排除/user/xx.txt ,则创建user目录
func Mkdir(path string) {
	path = GetFilePath(path)
	//log.Println(path) //创建目录
	err := os.Mkdir(path, 0750)
	if os.IsExist(err) {
		return
	}
	if err != nil {
		log.Println(err.Error())
	}
}

// GetFilePath 返回一个目录格式 /user/xx.txt => /user/
func GetFilePath(path string) string {
	index := strings.LastIndex(path, GetPathSeparator())
	if index == -1 {
		panic("路径文件不对：" + path)
	}
	return path[:index]
}

// GetFilePathName 返回一个文件名 /user/xx.txt => xx.txt or xx.txt => xx.txt
func GetFilePathName(path string) string {
	index := strings.LastIndex(path, GetPathSeparator())
	if index == -1 {
		return path
	}
	return path[index:]
}

func ListFile(path string) []FileInfo {
	list := make([]FileInfo, 0)
	dir, err := os.ReadDir(path)
	if err != nil {
		if os.IsNotExist(err) {
			panic("目录不存在！")
		}
		panic(err)
	}
	for _, item := range dir {
		f := FileInfo{FileName: item.Name(), IsDirectory: 0}
		if item.IsDir() {
			f.FileName += GetPathSeparator()
			f.IsDirectory = 1
		}
		list = append(list, f)
	}
	return list
}

type FileInfo struct {
	// 文件名
	FileName string `from:"fileName" json:"fileName"`
	// 1 是目录，0 默认文件
	IsDirectory int8 `from:"isDirectory" json:"isDirectory"`
}

func CreateFile(path string) *os.File {
	Mkdir(path)
	file := OpenFile(path)
	defer file.Close()
	return file
}

// OpenFile open File
func OpenFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

// ReadFile read File
func ReadFile(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return file
}

// FileMerge 指定合并某个目录下的文件,合并文件
func FileMerge(fileName, targetPath, timestamp, path string) {
	//最终文件路径
	var filePath = path + targetPath + fileName
	//分块目录路径
	blockPath := path + Md5Encode(fileName) + GetPathSeparator()
	Mkdir(filePath)
	file := CreateFile(filePath)
	files := ListFile(blockPath)
	for _, info := range files {
		var v = blockPath + info.FileName
		//log.Println(v) 打印合并路径
		bytes, err := os.ReadFile(v)
		if err != nil {
			panic(err)
		}
		_, err = file.Write(bytes)
		if err != nil {
			panic(err)
		}
		_ = os.Remove(v)
	}
	//解析10位时间戳并转换位当前日期对象
	time := ParseTime(timestamp)
	_ = os.Chtimes(filePath, time, time)
	defer file.Close()
	_ = os.Remove(blockPath)
}

// GetPathSeparator 获取系统路径分割符号 linux = / or win =\\
func GetPathSeparator() string {
	return pathSeparator
}
