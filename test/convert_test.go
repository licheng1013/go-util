package test

import (
	"gitee.com/licheng1013/go-util/common"
	"testing"
)

func TestInt(t *testing.T) {
	t.Log(common.ConvertUtil.ToInt("11"))
	t.Log(common.ConvertUtil.ToString(11))
}
