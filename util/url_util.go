package util

import (
	"net/url"
)

// UrlUtil url编码工具
type UrlUtil struct {
}

func NewUrlUtil() *UrlUtil {
	return &UrlUtil{}
}

// URLEncode 编码网址
func (UrlUtil) URLEncode(str string) string {
	return url.QueryEscape(str)
}

// URLDecode 解码网址
func (UrlUtil) URLDecode(str string) string {
	s, err := url.QueryUnescape(str)
	if err != nil {
		panic(err)
	}
	return s
}
