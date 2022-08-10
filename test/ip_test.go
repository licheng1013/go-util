/*
 * 作者：Lc
 */

package test

import (
	"gitee.com/licheng1013/go-util/util"
	"testing"
)

func TestIp(t *testing.T) {
	ip := util.GetIp()
	t.Log(ip)
	maskIps := util.GetTargetMaskIp(24)
	t.Log(maskIps)
	for _, maskIp := range maskIps {
		t.Log("http://" + maskIp)
	}
}
