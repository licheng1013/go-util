/*
 * 作者：Lc
 */
package test

import (
	"gitee.com/licheng1013/go-util/common"
	"log"
	"testing"
)

var filePath = "./test/test.txt"
var fileDirectory = "test"

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
	common.FileUtil.OpenFile(fileDirectory)
}
