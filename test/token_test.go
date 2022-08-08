/*
 * 作者：Lc
 */

package test

import (
	"go-util/util"
	"testing"
)

func TestToken(t *testing.T) {
	token := util.GetToken("12345")
	t.Log(token)
	t.Log(util.GetUserId(token))
}
