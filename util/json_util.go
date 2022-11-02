package util

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

// JsonUtil json转换工具类
type JsonUtil struct {
}

func NewJsonUtil() *JsonUtil {
	return &JsonUtil{}
}

// JsonToStr 把对象转换Json字符串！
func (JsonUtil) JsonToStr(v interface{}) string {
	if bytes, err := json.Marshal(v); err != nil {
		panic(err)
	} else {
		return string(bytes)
	}
}

// JsonToMap Json转换成对象 => 传参需要带&号 示例=> (jsonStr,&map)
func (JsonUtil) JsonToMap(jsonStr string, v interface{}) {
	if err := json.Unmarshal([]byte(jsonStr), v); err != nil {
		panic(err)
	}
}

// MapTosStruct map转换为结构体，需要注意的是，转换的字段必须大写开头
func (JsonUtil) MapTosStruct(in interface{}, out interface{}) {
	err := mapstructure.Decode(in, out)
	if err != nil {
		panic(err)
	}
}
