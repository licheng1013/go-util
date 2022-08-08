package util

import (
	"net/url"
)

// URLEncode 编码网址
func URLEncode(str string) string {
	return url.QueryEscape(str)
}

// URLDecode 解码网址
func URLDecode(str string) string {
	s, err := url.QueryUnescape(str)
	if err != nil {
		panic(err)
	}
	return s
}
