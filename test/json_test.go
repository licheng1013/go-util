/*
 * 作者：Lc
 */

package test

import (
	"gitee.com/licheng1013/go-util/util"
	"testing"
)

type Hello struct {
	Name string
	Age  int
}

func TestJson(t *testing.T) {
	var m = map[string]string{"A": "b"}
	str := util.JsonToStr(m)
	t.Log(str)
	var v map[string]string
	util.JsonToObj(str, &v)
	t.Log(v)
}

func TestStruct(t *testing.T) {
	m := map[string]interface{}{"age": 1, "name": "hi"}
	var v Hello
	util.MapToObj(m, &v)
	t.Log(v)
}
