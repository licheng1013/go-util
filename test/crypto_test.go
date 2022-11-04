package test

import (
	"gitee.com/licheng1013/go-util/common"
	"log"
	"testing"
)

func TestMD5(t *testing.T) {
	encode := common.CryptoUtil.Md5Encode("123456")
	log.Println(encode == "e10adc3949ba59abbe56e057f20f883e")
}

func TestRSA(t *testing.T) {
	rsa := common.CryptoUtil.RsaCreate()
	log.Println(rsa)
}
