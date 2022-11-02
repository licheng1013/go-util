package util

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

// JsonToStr 把对象转换Json字符串！
func JsonToStr(v interface{}) string {
	if bytes, err := json.Marshal(v); err != nil {
		panic(err)
	} else {
		return string(bytes)
	}
}

// JsonToObj Json转换成对象 => 传参需要带&号 示例=> (jsonStr,&user)
func JsonToObj(jsonStr string, v interface{}) {
	if err := json.Unmarshal([]byte(jsonStr), v); err != nil {
		panic(err)
	}
}

// MapToObj map转换为结构体，需要注意的是，转换的字段必须大写开头
func MapToObj(in interface{}, out interface{}) {
	err := mapstructure.Decode(in, out)
	if err != nil {
		panic(err)
	}
}
