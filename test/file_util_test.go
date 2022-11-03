package test

import (
	"gitee.com/licheng1013/go-util/common"
	"log"
	"testing"
)

var filePath = "./test/test.txt"
var fileDirectory = "./test"

func TestIsFile(t *testing.T) {
	file, err := common.FileUtil.IsFile(filePath)
	if err != nil {
		panic(err)
	}
	log.Printf("是否是文件: %v - 预期结果: true", file)
}

func TestIsDirectory(t *testing.T) {
	file, err := common.FileUtil.IsDirectory(fileDirectory)
	if err != nil {
		panic(err)
	}
	log.Printf("是否是目录: %v - 预期结果: true", file)
}

func TestIsAbsolute(t *testing.T) {
	absolute, err := common.FileUtil.IsAbsolute(filePath)
	if err != nil {
		panic(err)
	}
	log.Println("绝对路径: ", absolute)

}

func TestRedaFile(t *testing.T) {
	body, err := common.FileUtil.RedaFile(filePath)
	if err != nil {
		panic(err)
	}
	log.Println("文件内容: ", body)
}

func TestOpenFile(t *testing.T) {
	edit := common.FileUtil.OpenFile(filePath)
	edit.WriteString("Hello")
	edit.WriteStringLn("Hello")
	edit.Close()
}

func TestOpenNewFile(t *testing.T) {
	edit := common.FileUtil.OpenNewFile(filePath)
	edit.WriteStringLn("Hello")
	edit.WriteString("Hello")
	edit.Close()
}

func TestCreateDirectory(t *testing.T) {
	var testPath = fileDirectory + "/ok/test"
	err := common.FileUtil.CreateDirectory(testPath)
	if err != nil {
		panic(err)
	}
}

func TestDeleteDirectory(t *testing.T) {
	var testPath = fileDirectory + "/ok"
	_ = common.FileUtil.DeleteAllFileOrDirectory(testPath)
}

func TestListFile(t *testing.T) {
	var testPath = fileDirectory
	v := common.FileUtil.ListFile(testPath)
	log.Println(v)
}
