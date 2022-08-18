/*
 * 作者：Lc
 */

package test

import (
	"gitee.com/licheng1013/go-util/util"
	"testing"
)

func TestTime(t *testing.T) {
	t.Log("解析时间")
	t.Log(util.ParseTimeStr("1660105962"))

	t.Log("获取10位时间戳和毫秒时间戳")
	t.Log(util.GetTodayTimestamp())
	t.Log(util.GetTodayMillisecondTimestamp())

	t.Log("今天时间")
	t.Log(util.GetTodayDate())
	t.Log(util.GetTodayDateTime())

	t.Log("今天开始结束时间")
	t.Log(util.GetTodayStartDate())
	t.Log(util.GetTodayEndDate())

	t.Log("月开始结束时间")
	t.Log(util.GetMonthStartDate())
	t.Log(util.GetMonthEndDate())

	t.Log("年开始结束时间")
	t.Log(util.GetYearStartDate())
	t.Log(util.GetYearEndDate())

}
