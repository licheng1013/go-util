package test

import (
	"gitee.com/licheng1013/go-util/util"
	"testing"
)

func TestToken(t *testing.T) {
	var TokenUtil = util.NewTokenUtil("asdfas@dafdsfa@dfa")

	token := TokenUtil.DefaultToken("12345")
	t.Log(token)
	t.Log(TokenUtil.ParseToken(token))
}
