package test

import (
	"gitee.com/licheng1013/go-util/common"
	"log"
	"testing"
)

func TestDownload(t *testing.T) {
	url := "http://192.168.56.1:10200/file/download?path=E:\\my-study\\file-upload\\api\\Godot_v3.5.1-stable_win64.exe&name=app.go"
	common.HttpUtil.Download(url, "app.go")
}
func TestGet(t *testing.T) {
	url := "http://192.168.56.1:10200/file/list"
	get := common.HttpUtil.Get(url)
	log.Println(get)
}
