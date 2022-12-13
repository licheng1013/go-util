package util

import "strconv"

type ConvertUtil struct {
}

func NewConvertUtil() *ConvertUtil {
	return &ConvertUtil{}
}

// ToInt 字符串转Int64
func (c ConvertUtil) ToInt(str string) (number int64, err error) {
	number, err = strconv.ParseInt(str, 10, 64)
	return
}

// ToString 数字转String
func (c ConvertUtil) ToString(num int64) (str string) {
	str = strconv.FormatInt(num, 10)
	return
}
