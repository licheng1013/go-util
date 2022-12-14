package util

import (
	"gitee.com/licheng1013/go-util/model"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var pathSeparator = string(os.PathSeparator)

// FileUtil 文件工具类
type FileUtil struct {
}

func NewFileUtil() *FileUtil {
	return &FileUtil{}
}

// RedaFile 读取文件
func (v FileUtil) RedaFile(path string) (string, error) {
	absolute := v.GetAbsolute(path)
	file, err := os.ReadFile(absolute)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

// IsFile 是否是文件
func (v FileUtil) IsFile(path string) (bool, error) {
	info, err := v.IsDirectory(v.GetAbsolute(path))
	if err != nil { // 如果出现错误
		return false, err
	}
	return !info, err
}

// IsDirectory 是否是目录
func (v FileUtil) IsDirectory(path string) (bool, error) {
	info, err := os.Stat(v.GetAbsolute(path))
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

// IsAbsolute 是否绝对路径
func (v FileUtil) IsAbsolute(path string) (string, error) {
	abs, err := filepath.Abs(path)
	return abs, err
}

// GetAbsolute 从当前目录获取绝对路径
func (v FileUtil) GetAbsolute(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return abs
}

// OpenFile 打开文件,不存在则创建，以追加方式添加字符串
func (v FileUtil) OpenFile(path string) *model.FileEdit {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return &model.FileEdit{File: file, Path: path, Name: v.GetFileName(path)}
}

// OpenNewFile 打开新文件,与 OpenFile 不同的是，是先尝试删除后创建一个新文件在打开
func (v FileUtil) OpenNewFile(path string) *model.FileEdit {
	_ = v.DeleteFileOrDirectory(path)
	return v.OpenFile(path)
}

// CreateDirectory 创建目录 -> 支持子目录创建
func (v FileUtil) CreateDirectory(path string) error {
	err := os.MkdirAll(path, 0750)
	return err
}

// DeleteFileOrDirectory 删除文件或目录 -> 不包括子目录
func (v FileUtil) DeleteFileOrDirectory(path string) error {
	err := os.Remove(path)
	return err
}

// DeleteAllFileOrDirectory 删除文件或目录 -> 包括子目录
func (v FileUtil) DeleteAllFileOrDirectory(path string) error {
	err := os.RemoveAll(path)
	return err
}

// GetFileName 截取路径返回一个文件名 /user/xx.txt -> xx.txt ,如果没有符合的直接返回改字符串
func (v FileUtil) GetFileName(path string) string {
	index := strings.LastIndex(path, v.PathSeparator())
	if index == -1 {
		return path
	}
	return path[index:]
}

// FileMerge 指定合并某个目录下的文件,合并文件
//func (v FileUtil) FileMerge(fileName, targetPath, timestamp, path string) {
//	//最终文件路径
//	var filePath = path + targetPath + fileName
//	//分块目录路径
//	var cUtil CryptoUtil
//	blockPath := path + cUtil.Md5Encode(fileName) + v.PathSeparator()
//	_ = v.CreateDirectory(filePath)
//	file := v.OpenFile(filePath).File
//	files := v.ListFileAndFolder(blockPath)
//	for _, info := range files {
//		var v = blockPath + info.FileName
//		//log.Println(v) 打印合并路径
//		bytes, err := os.ReadFile(v)
//		if err != nil {
//			panic(err)
//		}
//		_, err = file.Write(bytes)
//		if err != nil {
//			panic(err)
//		}
//		_ = os.Remove(v)
//	}
//	//解析10位时间戳并转换位当前日期对象
//	var timeUtil TimeUtil
//	time := timeUtil.ParseTime(timestamp)
//	_ = os.Chtimes(filePath, time, time)
//	defer file.Close()
//	_ = os.Remove(blockPath)
//}

// PathSeparator 获取系统路径分割符号 linux = / or win =\\
func (FileUtil) PathSeparator() string {
	return pathSeparator
}

// ListFileAndFolder 列出指定目录文件不包括子目录 - 如果有错误会直接抛出
func (v FileUtil) ListFileAndFolder(path string) []model.FileInfo {
	path = v.GetAbsolute(path)
	separator := v.PathSeparator()
	list := make([]model.FileInfo, 0)
	dir, err := os.ReadDir(path)
	if err != nil {
		if os.IsNotExist(err) {
			panic("目录不存在！")
		}
		panic(err)
	}
	for _, item := range dir {
		info, _ := item.Info()
		f := model.FileInfo{FileName: item.Name(), IsDirectory: 0,
			FilePath: path + separator + item.Name(), UpdatedAt: info.ModTime().Format("2006-01-02 15:04:05")}

		if item.IsDir() {
			f.FilePath += separator
			f.IsDirectory = 1
		}
		list = append(list, f)
	}
	return list
}

// ListAllFileAndFolder 列出所有文件包含目录和文件
func (v FileUtil) ListAllFileAndFolder(path string) []model.FileInfo {
	file := v.ListFileAndFolder(path)
	for _, info := range file {
		if info.IsDirectory == 1 {
			listFile := v.ListAllFileAndFolder(info.FilePath)
			for _, k := range listFile {
				file = append(file, k)
			}
		}
	}
	return file
}

// ListAllFile 列出所有文件包获子目录下的文件，不包含目录
func (v FileUtil) ListAllFile(path string) []model.FileInfo {
	var list []model.FileInfo
	folder := v.ListAllFileAndFolder(path)
	for _, info := range folder {
		if info.IsDirectory == 0 {
			list = append(list, info)
		}
	}
	return list
}

// WriteFile 根据流写入文件
func (v FileUtil) WriteFile(file io.Reader, path string) error {
	all, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	edit := v.OpenNewFile(path)
	edit.WriteByte(all)
	edit.Close()
	return nil
}
