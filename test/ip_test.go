/*
 * 作者：Lc
 */

package test

import (
	"gitee.com/licheng1013/go-util/common"
	"testing"
)

func TestIp(t *testing.T) {
	ip := common.IpUtil.GetIp()
	t.Log(ip)
	maskIps := common.IpUtil.GetTargetMaskIp(24)
	t.Log(maskIps)
	for _, maskIp := range maskIps {
		t.Log("http://" + maskIp)
	}
}
