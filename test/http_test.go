package test

import (
	"gitee.com/licheng1013/go-util/common"
	"log"
	"testing"
)

func TestDownload(t *testing.T) {
	download := common.HttpUtil.Download("http://192.168.56.1:10200/file/download?path=E:\\my-study\\file-upload\\api\\app.go&name=app.go")
	log.Println(download)
}
