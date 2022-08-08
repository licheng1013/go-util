package util

import (
	"crypto/md5"
	"fmt"
)

func init() {
	//log.Println(md5Encode(123456))
}

// Md5Encode md5编码只有加密
func Md5Encode(v interface{}) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v", v))))
}
