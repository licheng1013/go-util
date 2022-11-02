package test

import (
	"gitee.com/licheng1013/go-util/util"
	"testing"
)

func TestUrl(t *testing.T) {
	encode := util.URLEncode("你好")
	t.Log(encode)
	t.Log(util.URLDecode(encode))
}
