/*
 * 作者：Lc
 */

package test

import (
	"go-util/util"
	"testing"
)

func TestTime(t *testing.T) {
	date := util.GetTodayDate()
	if len(date) != 10 {
		t.Error(date + " 日期错误!")
	}
	time := util.GetTodayDateTime()
	if len(time) != 19 {
		t.Error(time + " 日期错误!")
	}
}
