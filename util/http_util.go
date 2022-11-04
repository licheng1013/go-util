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

func (u HttpUtil) Download(url string, filePath string) {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	length := response.ContentLength
	log.Println("读取长度: ", length)
	util := NewFileUtil()
	file := util.OpenNewFile(util.GetAbsolute(filePath))
	for true {
		bytes := make([]byte, 1024*8)
		read, err := response.Body.Read(bytes)
		log.Println(read)
		if read == 0 && err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
		file.WriteByte(bytes[0:read])
	}
	file.Close()
}
