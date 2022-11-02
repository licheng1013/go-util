package test

import (
	"gitee.com/licheng1013/go-util/util"
	"testing"
)

func TestTime(t *testing.T) {
	var u = util.TimeUtil{}
	t.Log("解析时间")
	t.Log(u.ParseTimeStr("1660105962"))

	t.Log("获取10位时间戳和毫秒时间戳")
	t.Log(u.GetTodayTimestamp())
	t.Log(u.GetTodayMillisecondTimestamp())

	t.Log("今天时间")
	t.Log(u.GetTodayDate())
	t.Log(u.GetTodayDateTime())

	t.Log("今天开始结束时间")
	t.Log(u.GetTodayStartDate())
	t.Log(u.GetTodayEndDate())

	t.Log("月开始结束时间")
	t.Log(u.GetMonthStartDate())
	t.Log(u.GetMonthEndDate())

	t.Log("年开始结束时间")
	t.Log(u.GetYearStartDate())
	t.Log(u.GetYearEndDate())

}
