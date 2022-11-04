package test

import (
	"gitee.com/licheng1013/go-util/common"
	"testing"
)

func TestUrl(t *testing.T) {
	encode := common.UrlUtil.URLEncode("你好")
	t.Log(encode)
	b := common.UrlUtil.URLDecode(encode) == "你好"
	if !b {
		panic("解码失败！")
	}
}
