package util

import (
	"context"
	"github.com/gin-gonic/gin"
	"strings"
)

// GinUtil 对gin框架进行的简易封装
type GinUtil struct {
	ctx context.Context
	ket string
}

func NewGinUtil() *GinUtil {
	return &GinUtil{ket: "ginUtilKey", ctx: context.Background()}
}

// GetUrlPath 获取不带参数的路径  /a/b?name=xx -> /a/b
func (v GinUtil) GetUrlPath() string {
	uri := v.GetRequestURI()
	split := strings.Split(v.GetRequestURI(), "?")
	if len(split) >= 1 {
		return split[0]
	}
	return uri
}

// GetRequestURI 获取请求路径
func (v GinUtil) GetRequestURI() string {
	c := v.GetContext()
	return c.Request.RequestURI
}

// GetContext 获取 gin 上下文
func (v GinUtil) GetContext() *gin.Context {
	value := v.ctx.Value(v.ket)
	if value == nil {
		panic("没有在gin环境注册！")
	}
	return value.(*gin.Context)
}

// SetContext 设置 gin 上下文
func (v GinUtil) SetContext(c *gin.Context) {
	v.ctx = context.WithValue(v.ctx, v.ket, c)
}
