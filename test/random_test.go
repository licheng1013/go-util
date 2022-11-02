package test

import (
	"gitee.com/licheng1013/go-util/common"
	"testing"
)

func TestRandom(t *testing.T) {
	t.Log(common.RandomUtil.RandomNumber(10))
	t.Log(common.RandomUtil.RandomString(26))
	t.Log(common.RandomUtil.RandomRangeNumPlus(10, 30))
}
