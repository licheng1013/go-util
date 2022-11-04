package util

import (
	"io"
	"log"
	"net/http"
)

// HttpUtil http请求工具类
type HttpUtil struct {
}

func NewHttpUtil() *HttpUtil {
	return &HttpUtil{}
}

func (u HttpUtil) Get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	all, _ := io.ReadAll(response.Body)
	return string(all)
}

func (u HttpUtil) Download(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	all, _ := io.ReadAll(response.Body)
	return string(all)
}
