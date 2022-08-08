package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TimeFormat 格式必须如下1个字符也不要动！
const TimeFormat = "2006-01-02 15:04:05"

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

// GetTodayTime 获取当前时间10位时间戳,秒
func GetTodayTime() int64 {
	return time.Now().Unix()
}

// GetTodayTimeMillisecond 获取当前时间时间戳
func GetTodayTimeMillisecond() int64 {
	return time.Now().UnixMilli()
}

// GetTodayDateTime 获取当前日期时间  2006-01-02 15:04:05
func GetTodayDateTime() string {
	return ParseTimeStr(GetTodayTime())
}

// GetTodayDate 获取当前日期 2006-01-02
func GetTodayDate() string {
	return ParseTimeStrAndFormatStr(GetTodayTime(), strings.Split(TimeFormat, " ")[0])
}
