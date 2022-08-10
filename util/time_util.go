package util

import (
	"fmt"
	"strconv"
	"time"
)

// TimeFormat 格式必须如下1个字符也不要动！
const TimeFormat = YearMonthDayDateTime

//格式必须如下1个字符也不要动！

// YearMonthDayDateTime 年月日
const YearMonthDayDateTime = "2006-01-02 15:04:05"
const YearMonthDay = "2006-01-02"
const YearMonth = "2006-01"
const Year = "2006"

const EndTime = " 23:59:59"
const StartTime = " 00:00:00"

// ParseTimeStr 解析10位时间戳,返回格式好的日期格式 2006-01-02 15:04:05
func ParseTimeStr(t interface{}) string {
	return ParseTimeStrAndFormatStr(t, TimeFormat)
}

// ParseTime 解析10位时间戳,返回日期对象
func ParseTime(t string) time.Time {
	v, err := strconv.ParseInt(t, 10, 64)
	if len(t) != 10 {
		panic("不是10位时间戳！")
	}
	if err != nil {
		panic(err)
	}
	return time.Unix(v, 0)
}

// ParseTimeStrAndFormatStr 解析10位时间戳
func ParseTimeStrAndFormatStr(t interface{}, f string) string {
	str := fmt.Sprintf("%v", t)
	parseInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	format := time.Unix(parseInt, 0).Format(f)
	return format
}

// GetTodayTimestamp 获取当前时间10位时间戳,秒
func GetTodayTimestamp() int64 {
	return time.Now().Unix()
}

// GetTodayMillisecondTimestamp 获取当前时间13位时间戳,毫秒
func GetTodayMillisecondTimestamp() int64 {
	return time.Now().UnixMilli()
}

// GetTodayDateTime 获取当前日期时间  2006-01-02 15:04:05
func GetTodayDateTime() string {
	return time.Now().Format(YearMonthDayDateTime)
}

// GetTodayDate 获取当前日期 2006-01-02
func GetTodayDate() string {
	return time.Now().Format(YearMonthDay)
}

// GetTodayStartDate 获取当前日期 2006-01-02 00:00:00
func GetTodayStartDate() string {
	return time.Now().Format(YearMonthDay) + StartTime
}

// GetTodayEndDate 获取当前日期 2006-01-02 23:59:59
func GetTodayEndDate() string {
	return time.Now().Format(YearMonthDay) + EndTime
}

// GetMonthStartDate 获取当前月开始日期 示例: 2006-02-12 => 2006-02-01
func GetMonthStartDate() string {
	now := time.Now()
	now = now.AddDate(0, 0, -now.Day()+1)
	return now.Format(YearMonthDay) + StartTime
}

// GetMonthEndDate 获取当前月结束日期 示例: 2006-02-12 => 2006-02-30
func GetMonthEndDate() string {
	now := time.Now()
	now = now.AddDate(0, +1, -now.Day())
	return now.Format(YearMonthDay) + EndTime
}

// GetYearStartDate 获取当前年开始日期 示例: 2006-02-12 => 2006-01-01
func GetYearStartDate() string {
	now := time.Now()
	return now.Format(Year) + "-01-01" + StartTime
}

// GetYearEndDate 获取当前年结束日期 示例: 2006-02-12 => 2006-12-31
func GetYearEndDate() string {
	now := time.Now()
	return now.Format(Year) + "-12-31" + EndTime
}
