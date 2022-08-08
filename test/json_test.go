/*
 * 作者：Lc
 */

package test

import (
	"go-util/util"
	"testing"
)

func TestJson(t *testing.T) {
	var m = map[string]string{"A": "b"}
	str := util.JsonToStr(m)
	t.Log(str)
	var v map[string]string
	util.JsonToObj(str, &v)
	t.Log(v)
}
